const API = "https://YOUR_RENDER_URL";

function updateProgress(data) {
  fetch(`${API}/api/progress/update`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Authorization": "Bearer " + localStorage.getItem("token")
    },
    body: JSON.stringify(data)
  });
}
