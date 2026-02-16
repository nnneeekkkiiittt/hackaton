const API_BASE = "http://localhost:8080";

document.addEventListener("DOMContentLoaded", function () {
    const params = new URLSearchParams(window.location.search);
    const id = params.get("id");
    const nameEl = document.getElementById("buildingName");
    const descEl = document.getElementById("buildingDescription");
    const mapEl = document.getElementById("buildingMap");
    const errEl = document.getElementById("buildingError");

    if (!id) {
        if (errEl) errEl.textContent = "Не указано здание. Вернитесь на маршрут и выберите здание.";
        return;
    }

    fetch(`${API_BASE}/building/${id}`)
        .then((response) => {
            if (!response.ok) throw new Error("Not found");
            return response.json();
        })
        .then((building) => {
            if (nameEl) nameEl.textContent = building.name || "Здание";
            document.title = (building.name || "Здание") + " — ФинКод Нижнего";

            // Description from API endpoint (backend reads .txt from pkg)
            fetch(`${API_BASE}/building/${id}/description`)
                .then((r) => (r.ok ? r.text() : Promise.reject()))
                .then((text) => {
                    if (descEl) {
                        descEl.textContent = text;
                        descEl.classList.remove("building-description-empty");
                    }
                })
                .catch(() => {
                    if (descEl) {
                        descEl.textContent = building.description || "";
                        if (!building.description) descEl.classList.add("building-description-empty");
                    }
                });

            // Yandex Map: show building location if we have coordinates
            const lat = building.latitude;
            const lon = building.longitude;
            if (mapEl && typeof lat === "number" && typeof lon === "number" && !Number.isNaN(lat) && !Number.isNaN(lon)) {
                if (typeof ymaps !== "undefined") {
                    ymaps.ready(function () {
                        const map = new ymaps.Map("buildingMap", {
                            center: [lat, lon],
                            zoom: 16,
                            controls: ["zoomControl", "typeSelector", "fullscreenControl"]
                        });
                        const placemark = new ymaps.Placemark(
                            [lat, lon],
                            {
                                balloonContent: building.name || "Здание",
                                hintContent: building.name || "Здание"
                            },
                            {
                                preset: "islands#redIcon"
                            }
                        );
                        map.geoObjects.add(placemark);
                    });
                }
            } else if (mapEl) {
                mapEl.style.display = "none";
            }
        })
        .catch(() => {
            if (errEl) errEl.textContent = "Здание не найдено или произошла ошибка загрузки.";
        });
});
