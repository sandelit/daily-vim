const tips = [
  "Use 'dd' to delete a line",
  "Press 'i' to enter insert mode",
  "Use '/pattern' to search for a pattern",
  // Add more tips here
];

function displayRandomTip() {
  const tipElement = document.getElementById("tip");
  const randomIndex = Math.floor(Math.random() * tips.length);
  tipElement.textContent = tips[randomIndex];
}

// Display a tip when the page loads
window.onload = displayRandomTip;
