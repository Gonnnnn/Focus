const COMPLETE = "COMPLETE";
const IN_PROGRESS = "IN_PROGRESS";
const NOT_STARTED = "NOT_STARTED";
const EXPIRED = "EXPIRED";
const BASE_URL = "http://localhost:8080";

function initActivityItem(activityItem) {
  // Initialize timestamps.
  const currentTime = new Date();
  const startTimestampElement = activityItem.querySelector(
    ".activity-start-timestamp-value"
  );
  const startTimestampMilli = parseInt(
    startTimestampElement.textContent.trim()
  );
  const endTimestampElement = activityItem.querySelector(
    ".activity-end-timestamp-value"
  );
  const endTimestampMilli = parseInt(endTimestampElement.textContent.trim());
  startTimestampElement.textContent = formatKoreanTime(
    new Date(startTimestampMilli)
  );
  endTimestampElement.textContent = formatKoreanTime(
    new Date(endTimestampMilli)
  );

  // Initialize status.
  const statusValueElement = activityItem.querySelector(
    ".activity-status-value"
  );
  const status = statusValueElement.textContent.trim();
  statusValueElement.style.color = statusColor(status);

  // Initialize progress bar.
  const progressPercentage = calculateProgress(
    currentTime.getTime(),
    startTimestampMilli,
    endTimestampMilli
  );
  const progressBar = document.createElement("div");
  progressBar.className = "activity-progress-bar";
  progressBar.textContent = progressText(progressPercentage);
  activityItem.querySelector(".activity-item-content").appendChild(progressBar);
  if (status === IN_PROGRESS) {
    setInterval(function () {
      const progressPercentage = calculateProgress(
        new Date().getTime(),
        startTimestampMilli,
        endTimestampMilli
      );

      progressBar.textContent = progressText(progressPercentage);
      if (progressPercentage >= 100) {
        // Reload page when activity is completed. It can cause a slight problem when
        // there are multiple activities in the same page, but it's not a big deal.
        location.reload();
      }
    }, 60000);
  }

  // Attach event listeners to buttons.
  const id = activityItem.getAttribute("activity-id");
  const deleteButton = activityItem.querySelector(".delete-button");
  deleteButton.addEventListener("click", (_) => deleteButtonClickHandler(id));
  const completeButton = activityItem.querySelector(".complete-button");
  if (completeButton !== null) {
    completeButton.addEventListener("click", (_) =>
      completeButtonClickHandler(id)
    );
  }
}

function initCreateForm(createActivityForm) {
  const startTimeInput = createActivityForm.querySelector("#startTime");
  const endTimeInput = createActivityForm.querySelector("#endTime");

  const defaultTimeInput = getCreateFormDateTime(new Date());
  startTimeInput.value = defaultTimeInput;
  endTimeInput.value = defaultTimeInput;

  createActivityForm.addEventListener("submit", createFormClickHandler);
}

async function deleteButtonClickHandler(id) {
  const confirmation = confirm(
    "Are you sure you want to delete this activity?"
  );

  if (!confirmation) return;

  const response = await deleteActivity(id);
  if (!response.ok) {
    const reason = await response.text();
    console.error(`Failed to delete activity. Reason: ${reason}`);
    alert(`Failed to delete activity: ${reason}.`);
    return;
  }

  location.reload();
}

async function completeButtonClickHandler(id) {
  const confirmation = confirm(
    "Are you sure you want to complete this activity?"
  );

  if (!confirmation) return;

  const response = await completeActivity(id);
  if (!response.ok) {
    const reason = await response.text();
    console.error(`Failed to complete activity. Reason: ${reason}`);
    alert(`Failed to complete activity: ${reason}.`);
    return;
  }

  location.reload();
}

async function createFormClickHandler(event) {
  event.preventDefault();

  const formData = new FormData(this);
  const activityName = formData.get("title");
  const activityDescription = formData.get("description");
  const startTimeMilli = new Date(formData.get("startTime")).getTime();
  const endTimeMilli = new Date(formData.get("endTime")).getTime();

  if (startTimeMilli >= endTimeMilli) {
    alert("Start time must be before end time.");
    return;
  }

  const response = await createActivity({
    title: activityName,
    description: activityDescription,
    startTimestampMilli: startTimeMilli,
    endTimestampMilli: endTimeMilli,
  });

  if (!response.ok) {
    const reason = await response.text();
    console.error(`Failed to create activity. Reason: ${reason}`);
    alert(`Failed to create activity: ${reason}.`);
    return;
  }

  location.reload();
}

async function createActivity(activity) {
  return await fetch(BASE_URL, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(activity),
  })
    .then((response) => {
      return response;
    })
    .catch((error) => {
      console.error("Network error:", error);
    });
}

async function deleteActivity(id) {
  const data = { id };
  return await fetch(BASE_URL, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  })
    .then((response) => {
      return response;
    })
    .catch((error) => {
      console.error("Network error:", error);
    });
}

async function completeActivity(id) {
  const data = { id };
  return await fetch(BASE_URL + "/complete", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  })
    .then((response) => {
      return response;
    })
    .catch((error) => {
      console.error("Network error:", error);
    });
}

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

function calculateProgress(
  currentTime,
  startTimestampMilli,
  endTimestampMilli
) {
  if (currentTime < startTimestampMilli) {
    return 0;
  } else if (currentTime >= endTimestampMilli) {
    return 100;
  } else {
    return (
      ((currentTime - startTimestampMilli) /
        (endTimestampMilli - startTimestampMilli)) *
      100
    );
  }
}

function progressText(percentage) {
  const barString =
    "■".repeat(Math.floor(percentage / 10)) +
    "□".repeat(Math.floor((100 - percentage) / 10));

  return `${barString} ${percentage.toFixed(2)}%`;
}

function getCreateFormDateTime(time) {
  const year = time.getFullYear();
  const month = String(time.getMonth() + 1).padStart(2, "0");
  const day = String(time.getDate()).padStart(2, "0");
  const hours = String(time.getHours()).padStart(2, "0");
  const minutes = String(time.getMinutes()).padStart(2, "0");

  return `${year}-${month}-${day}T${hours}:${minutes}`;
}

function statusColor(status) {
  switch (status) {
    case COMPLETE:
      return "#3CDE27"; // green
    case IN_PROGRESS:
      return "#1EB4D5"; // sky blue
    case NOT_STARTED:
      return "black";
    case EXPIRED:
      return "gray";
    default:
      return "black";
  }
}

document.addEventListener("DOMContentLoaded", function () {
  const activityItems = document.querySelectorAll(".activity-item");
  activityItems.forEach(initActivityItem);

  const createActivityForm = document.getElementById("createActivityForm");
  initCreateForm(createActivityForm);
});
