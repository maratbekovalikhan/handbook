const API = "https://YOUR_RENDER_URL";

async function loadProgress() {
  const token = localStorage.getItem("token");
  if (!token) {
    window.location.href = "login.html";
    return;
  }

  const res = await fetch(`${API}/api/progress/me`, {
    headers: {
      "Authorization": "Bearer " + token
    }
  });

  const data = await res.json();
  const container = document.getElementById("progress");

  container.innerHTML = "";
  data.forEach(p => {
    container.innerHTML += `
      <p>
        <b>${p.course.toUpperCase()}</b> — ${p.percent}%
      </p>
    `;
  });
}

loadProgress();
