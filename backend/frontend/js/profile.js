const API = "https://handbook-backend-ah0j.onrender.com";

async function loadProfile() {
  const token = localStorage.getItem("token");

  if (!token) {
    window.location.href = "login.html";
    return;
  }

  const res = await fetch(`${API}/api/profile`, {
    headers: {
      "Authorization": "Bearer " + token
    }
  });

  if (!res.ok) {
    localStorage.removeItem("token");
    window.location.href = "login.html";
    return;
  }

  const user = await res.json();

  document.getElementById("profile").innerHTML = `
    <p><strong>Имя:</strong> ${user.name}</p>
    <p><strong>Email:</strong> ${user.email}</p>
  `;
}

loadProfile();
