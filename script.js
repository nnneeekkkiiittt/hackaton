const API_BASE = "http://localhost:8080";

document.addEventListener("DOMContentLoaded", function () {
    // Log out: clear user when landing on index with #logout
    if (window.location.hash === "#logout") {
        localStorage.removeItem("user");
        window.location.hash = "";
    }

    const registrationForm = document.getElementById("registrationForm");
    const loginForm = document.getElementById("loginForm");

    // ----- Auth tabs (index page) -----
    const authTabs = document.querySelectorAll(".auth-tab");
    const loginPanel = document.getElementById("login-panel");
    const registerPanel = document.getElementById("register-panel");

    if (authTabs.length) {
        authTabs.forEach((tab) => {
            tab.addEventListener("click", function () {
                const target = this.getAttribute("data-tab");
                authTabs.forEach((t) => t.classList.remove("active"));
                this.classList.add("active");
                if (target === "login") {
                    loginPanel.classList.add("active");
                    registerPanel.classList.remove("active");
                } else {
                    registerPanel.classList.add("active");
                    loginPanel.classList.remove("active");
                }
            });
        });
    }

    // ----- Registration (index + register page) -----
    if (registrationForm) {
        registrationForm.addEventListener("submit", function (e) {
            e.preventDefault();
            const errEl = document.getElementById("registrationError");
            if (errEl) errEl.textContent = "";

            const username = document.getElementById("username").value.trim();
            const email = document.getElementById("email").value.trim();
            const password = document.getElementById("password").value;

            fetch(`${API_BASE}/user`, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ username, email, password }),
            })
                .then((response) => {
                    if (!response.ok) return response.json().then((d) => Promise.reject(d));
                    return response.json();
                })
                .then((data) => {
                    localStorage.setItem("user", JSON.stringify(data));
                    window.location.href = "profile.html";
                })
                .catch((err) => {
                    const msg = err && err.error ? err.error : "Registration failed. Try again.";
                    if (errEl) errEl.textContent = msg;
                });
        });
    }

    // ----- Login (index page; backend has no /login — use GET /users + match) -----
    if (loginForm) {
        loginForm.addEventListener("submit", function (e) {
            e.preventDefault();
            const errEl = document.getElementById("loginError");
            if (errEl) errEl.textContent = "";

            const email = document.getElementById("loginEmail").value.trim();
            const password = document.getElementById("loginPassword").value;

            fetch(`${API_BASE}/users`)
                .then((response) => {
                    if (!response.ok) throw new Error("Login failed");
                    return response.json();
                })
                .then((users) => {
                    const user = users.find((u) => u.email === email && u.password === password);
                    if (!user) {
                        if (errEl) errEl.textContent = "Invalid email or password.";
                        return;
                    }
                    localStorage.setItem("user", JSON.stringify(user));
                    window.location.href = "profile.html";
                })
                .catch(() => {
                    if (errEl) errEl.textContent = "Could not connect. Try again.";
                });
        });
    }

    // ----- Profile page only -----
    const userInfo = document.getElementById("userInfo");
    const buildingsContainer = document.getElementById("buildingsContainer");

    if (userInfo && buildingsContainer) {
        const user = JSON.parse(localStorage.getItem("user"));
        if (!user) {
            window.location.href = "index.html";
            return;
        }

        displayUserInfo(user);
        fetchUserProgress(user.id);
    }

    function displayUserInfo(user) {
        const nameEl = document.getElementById("userName");
        const emailEl = document.getElementById("userEmail");
        if (nameEl) nameEl.textContent = user.username || "User";
        if (emailEl) emailEl.textContent = user.email || "";
    }

    function fetchUserProgress(userId) {
        Promise.all([
            fetch(`${API_BASE}/progress/${userId}`).then((r) => r.json()),
            fetch(`${API_BASE}/buildings`).then((r) => r.json()),
        ])
            .then(([progressData, buildingsData]) => {
                displayBuildings(progressData || [], buildingsData || []);
            })
            .catch((error) => console.error("Error fetching progress or buildings:", error));
    }

    function displayBuildings(progressData, buildingsData) {
        if (!buildingsContainer) return;
        buildingsContainer.innerHTML = "";

        const buildingsList = (buildingsData || []).slice().sort((a, b) => a.id - b.id);
        const user = JSON.parse(localStorage.getItem("user"));
        const userId = user ? user.id : null;

        // Map: building_id -> progress (so we know which are visited)
        const progressByBuilding = (progressData || []).reduce((acc, p) => {
            acc[p.building_id] = p;
            return acc;
        }, {});

        // Show every building from the DB; use progress only to mark visited
        buildingsList.forEach((building, index) => {
            const progress = progressByBuilding[building.id];
            const isVisited = progress ? !!progress.flag : false;
            const name = building.name || `Building #${building.id}`;

            const card = document.createElement("div");
            card.className = "building-card";

            const step = document.createElement("span");
            step.className = "building-step";
            step.textContent = index + 1;

            const el = document.createElement("div");
            el.className = "building" + (isVisited ? " visited" : "");
            el.title = isVisited ? "Visited" : "Mark as visited";
            if (!isVisited && userId) {
                el.addEventListener("click", () => markBuildingAsVisited(building.id, userId));
            }

            const label = document.createElement("a");
            label.className = "building-name";
            label.href = `building.html?id=${building.id}`;
            label.textContent = name;

            card.appendChild(step);
            card.appendChild(el);
            card.appendChild(label);
            buildingsContainer.appendChild(card);

            // Route pointer to next stop (arrow between cards)
            if (index < buildingsList.length - 1) {
                const pointer = document.createElement("div");
                pointer.className = "route-pointer";
                pointer.setAttribute("aria-hidden", "true");
                buildingsContainer.appendChild(pointer);
            }
        });
    }

    function markBuildingAsVisited(buildingId, userId) {
        fetch(`${API_BASE}/progress/${userId}/${buildingId}`, {
            method: "PUT",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ flag: true }),
        })
            .then((response) => response.json())
            .then(() => {
                fetchUserProgress(userId);
            })
            .catch((error) => console.error("Error marking building as visited:", error));
    }
});
