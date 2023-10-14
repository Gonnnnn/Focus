document.addEventListener("DOMContentLoaded", function () {
  // Get all activity list items
  const activityItems = document.querySelectorAll(".activity-item");

  activityItems.forEach(function (activityItem) {
    const startTimestampElement = activityItem.querySelector(
      ".activity-start-timestamp-value"
    );
    const startTimestamp = parseInt(startTimestampElement.textContent.trim());

    const endTimestampElement = activityItem.querySelector(
      ".activity-end-timestamp-value"
    );
    const endTimestamp = parseInt(endTimestampElement.textContent.trim());

    // Format the start and end timestamps
    startTimestampElement.textContent = formatKoreanTime(
      new Date(startTimestamp / 1000000)
    );
    endTimestampElement.textContent = formatKoreanTime(
      new Date(endTimestamp / 1000000)
    );

    const currentTime = new Date().getTime() * 1000000;
    const progressPercentage = calculateProgress(
      currentTime,
      startTimestamp,
      endTimestamp
    );

    // Create and append a progress bar element
    const progressBar = createProgressBar(progressPercentage);

    // Append the progress bar to the activity item
    activityItem.appendChild(progressBar);
  });
});

function formatKoreanTime(date) {
  const options = {
    weekday: "short",
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
    timeZone: "Asia/Seoul",
  };

  return date.toLocaleDateString("ko-KR", options);
}

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
          window.location.reload();
        } else {
          alert("Failed to create activity. Please try again.");
        }
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  });

setInterval(function () {
  location.reload();
}, 60000);
