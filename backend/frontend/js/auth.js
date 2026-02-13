const API = "https://handbook-backend-ah0j.onrender.com";


document.addEventListener("DOMContentLoaded", function () {
  const token = localStorage.getItem("token");
  const authLinks = document.getElementById("authLinks");

  if (token && authLinks) {
    authLinks.innerHTML = `
      <a href="/profile.html">Profile</a>
      <a href="#" onclick="logout()">Logout</a>
    `;
  }
});

function logout() {
  localStorage.removeItem("token");
  alert("Вы вышли из аккаунта");
  window.location.href = "/index.html";
}

async function register() {
  const name = document.getElementById("name").value;
  const email = document.getElementById("email").value;
  const password = document.getElementById("password").value;

  const res = await fetch(`${API}/api/register`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ name, email, password })
  });

  if (res.ok) {
    alert("Регистрация успешна");
    window.location.href = "login.html";
  } else {
    alert("Ошибка регистрации");
  }
}

async function login() {
  const email = document.getElementById("email").value;
  const password = document.getElementById("password").value;

  const res = await fetch(`${API}/api/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password })
  });

  const data = await res.json();

  if (data.token) {
    localStorage.setItem("token", data.token);
    window.location.href = "profile.html";
  } else {
    alert("Неверные данные");
  }
}
