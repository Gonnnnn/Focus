document.addEventListener("DOMContentLoaded", function () {
  // Get all activity list items
  const activityItems = document.querySelectorAll(".activity-item");
  const currentTime = new Date();

  activityItems.forEach(function (activityItem) {
    const startTimestampElement = activityItem.querySelector(
      ".activity-start-timestamp-value"
    );
    const startTimestamp = parseInt(startTimestampElement.textContent.trim());

    const endTimestampElement = activityItem.querySelector(
      ".activity-end-timestamp-value"
    );
    const endTimestamp = parseInt(endTimestampElement.textContent.trim());

    const deleteButton = activityItem.querySelector(".delete-button");
    deleteButton.addEventListener("click", (event) => {
      const id = event.target.getAttribute("activity-id");
      const confirmation = confirm(
        "Are you sure you want to delete this activity?"
      );
      if (confirmation) {
        deleteActivity(id);
      }
    });

    const completeButton = activityItem.querySelector(".complete-button");
    if (completeButton !== null) {
      completeButton.addEventListener("click", (event) => {
        const id = event.target.getAttribute("activity-id");
        const confirmation = confirm(
          "Are you sure you want to complete this activity?"
        );
        if (confirmation) {
          completeActivity(id);
        }
      });
    }

    // Format the start and end timestamps
    startTimestampElement.textContent = formatKoreanTime(
      new Date(startTimestamp / 1000000)
    );
    endTimestampElement.textContent = formatKoreanTime(
      new Date(endTimestamp / 1000000)
    );

    const progressPercentage = calculateProgress(
      currentTime.getTime() * 1000000,
      startTimestamp,
      endTimestamp
    );

    // Create and append a progress bar element
    const progressBar = createProgressBar(progressPercentage);

    // Append the progress bar to the activity item
    activityItem
      .querySelector(".activity-item-content")
      .appendChild(progressBar);

    // Reload the page every minute to update the progress bar.
    setInterval(function () {
      location.reload();
    }, 60000);
  });

  const createActivityForm = document.getElementById("createActivityForm");
  const startTimeInput = createActivityForm.querySelector("#startTime");
  const endTimeInput = createActivityForm.querySelector("#endTime");

  // Set the start time to the current time
  const defaultTimeInput = getCreateFormDateTime(currentTime);
  startTimeInput.value = defaultTimeInput;
  endTimeInput.value = defaultTimeInput;

  // Set the create activity form handler.
  createActivityForm.addEventListener("submit", function (event) {
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

function deleteActivity(id) {
  const url = "http://localhost:8080";
  const data = { id };
  fetch(url, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  })
    .then((response) => {
      if (response.ok) {
        // Successfully deleted, you can update the UI here if needed
        console.log("Activity deleted successfully");
        window.location.reload();
      } else {
        // Handle error here
        console.error("Failed to delete activity");
      }
    })
    .catch((error) => {
      // Handle network error
      console.error("Network error:", error);
    });
}

function completeActivity(id) {
  const url = "http://localhost:8080/complete";
  const data = { id };
  fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  })
    .then(async (response) => {
      if (response.ok) {
        // Successfully deleted, you can update the UI here if needed
        console.log("Activity completed successfully");
        window.location.reload();
      } else {
        // Handle error here
        console.error("Failed to complete activity");
        const reason = await response.text();
        // convert response body to string
        alert(`Failed to complete activity: ${reason}.`);
      }
    })
    .catch((error) => {
      // Handle network error
      console.error("Network error:", error);
    });
}

function getCreateFormDateTime(time) {
  const year = time.getFullYear();
  const month = String(time.getMonth() + 1).padStart(2, "0");
  const day = String(time.getDate()).padStart(2, "0");
  const hours = String(time.getHours()).padStart(2, "0");
  const minutes = String(time.getMinutes()).padStart(2, "0");

  return `${year}-${month}-${day}T${hours}:${minutes}`;
}
