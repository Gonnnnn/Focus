document.addEventListener("DOMContentLoaded", function () {
  // Get all activity list items
  const activityItems = document.querySelectorAll(".activity-item");

  activityItems.forEach(function (activityItem) {
    const startTimestamp = parseInt(
      activityItem
        .querySelector(".activity-start-timestamp-value")
        .textContent.trim()
    );
    const endTimestamp = parseInt(
      activityItem
        .querySelector(".activity-end-timestamp-value")
        .textContent.trim()
    );

    const currentTime = new Date().getTime() * 1000000;
    const progressPercentage = calculateProgress(
      currentTime,
      startTimestamp,
      endTimestamp
    );
    console.log(progressPercentage);
    // Create and append a progress bar element
    const progressBar = createProgressBar(progressPercentage);

    // Append the progress bar to the activity item
    activityItem.appendChild(progressBar);
  });
});

// Calculate progress percentage
function calculateProgress(currentTime, startTimestamp, endTimestamp) {
  if (currentTime < startTimestamp) {
    return 0;
  } else if (currentTime >= endTimestamp) {
    return 100;
  } else {
    return (
      ((currentTime - startTimestamp) / (endTimestamp - startTimestamp)) * 100
    );
  }
}

// Create a progress bar element
function createProgressBar(percentage) {
  const progressBar = document.createElement("div");
  progressBar.className = "activity-progress-bar";

  // Create a string with □ and ■ characters based on the percentage
  const barString =
    "■".repeat(Math.floor(percentage / 10)) +
    "□".repeat(Math.floor((100 - percentage) / 10));

  progressBar.innerHTML = `<div class="activity-progress-text">${barString} ${percentage.toFixed(
    2
  )}%</div>`;

  return progressBar;
}

document
  .getElementById("createActivityForm")
  .addEventListener("submit", function (event) {
    event.preventDefault();

    const formData = new FormData(this);
    const activityName = formData.get("title");
    const activityDescription = formData.get("description");
    const startTimeNano =
      new Date(formData.get("startTime")).getTime() * 1000000;
    const endTimeNano = new Date(formData.get("endTime")).getTime() * 1000000;

    // Send a POST request to http://localhost:8080
    fetch("http://localhost:8080", {
      method: "POST",
      body: JSON.stringify({
        title: activityName,
        description: activityDescription,
        startTimestamp: startTimeNano,
        endTimestamp: endTimeNano,
      }),
    })
      .then((response) => {
        if (response.ok) {
          alert("Activity created successfully!");
        } else {
          alert("Failed to create activity. Please try again.");
        }
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  });
