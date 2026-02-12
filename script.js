document.addEventListener("DOMContentLoaded", function() {
    const registrationForm = document.getElementById("registrationForm");

    // Обработчик отправки формы регистрации
    registrationForm.addEventListener("submit", function(e) {
        e.preventDefault();  // Предотвращаем стандартное поведение формы

        // Получаем данные из формы
        const username = document.getElementById("username").value;
        const email = document.getElementById("email").value;
        const password = document.getElementById("password").value;

        // Формируем объект с данными пользователя
        const userData = {
            username: username,
            email: email,
            password: password
        };

        // Отправляем запрос на сервер для регистрации пользователя
        fetch('http://localhost:8080/user', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(userData)
        })
        .then(response => response.json())
        .then(data => {
            console.log("User created:", data);  // Логируем информацию о новом пользователе
            localStorage.setItem("user", JSON.stringify(data));  // Сохраняем пользователя в localStorage
            displayUserInfo(data);  // Отображаем информацию о пользователе
            fetchUserProgress(data.id);  // Загружаем прогресс для нового пользователя
        })
        .catch(error => {
            console.error('Error during registration:', error);
            alert("Error during registration");
        });
    });

    // Проверяем, есть ли пользователь в localStorage
    const user = JSON.parse(localStorage.getItem("user"));
    if (user) {
        displayUserInfo(user);  // Если есть, отображаем его данные
        fetchUserProgress(user.id);  // Загружаем прогресс
    } else {
        alert("Please register first");
    }
});

// Отображение информации о пользователе
function displayUserInfo(user) {
    document.getElementById("userName").textContent = `Username: ${user.username}`;
    document.getElementById("userEmail").textContent = `Email: ${user.email}`;
}

// Функция для получения прогресса пользователя (зданий)
function fetchUserProgress(userId) {
    fetch(`http://localhost:8080/progress/${userId}`)
    .then(response => response.json())
    .then(progressData => {
        displayBuildings(progressData);
    })
    .catch(error => console.error('Error fetching progress:', error));
}

// Отображение зданий на странице
function displayBuildings(progressData) {
    const buildingsContainer = document.getElementById("buildingsContainer");
    buildingsContainer.innerHTML = '';  // Очищаем контейнер

    progressData.forEach(progress => {
        const buildingElement = document.createElement('div');
        buildingElement.classList.add('building');
        
        // Если здание посещено, добавляем класс 'visited'
        if (progress.visited) {
            buildingElement.classList.add('visited');
        }

        // Добавляем обработчик для клика на здание
        buildingElement.addEventListener("click", () => markBuildingAsVisited(progress.building_id, progress.user_id));

        buildingsContainer.appendChild(buildingElement);
    });
}

// Отметить здание как посещенное
function markBuildingAsVisited(buildingId, userId) {
    fetch(`http://localhost:8080/progress/${userId}/${buildingId}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ visited: true })
    })
    .then(response => response.json())
    .then(data => {
        alert('Building marked as visited!');
        fetchUserProgress(userId);  // Обновляем прогресс
    })
    .catch(error => console.error('Error marking building as visited:', error));
}
