const temperatureData = {
  datasets: [
    {
      backgroundColor: "rgba(54, 162, 235, 0.4)",
      borderColor: "rgba(54, 162, 235, 1)",
      hoverBackgroundColor: "rgba(54, 162, 235, 0.9)",
      borderWidth: 2,
      borderSkipped: false,
      data: [
        { x: "Day 1", y: [-3, 5] }, 
        { x: "Day 2", y: [22, 32] }, 
        { x: "Day 3", y: [18, 28] }, 
        { x: "Day 4", y: [19, 29] }, 
        { x: "Day 5", y: [21, 31] }, 
      ],
    },
  ],
}

const ctx = document.getElementById("temperatureChart");

const toolTipTitle = (toolTipItems) => {
  return toolTipItems[0].label 
}

const toolTipLabel = (toolTipItem) => {
  return [`High: ${toolTipItem.raw.y[1]}`, `Low: ${toolTipItem.raw.y[0]}`]
}

const temperatureChart = new Chart(ctx, {
  type: "bar",
  data: temperatureData,
  options: {
    plugins: {
      legend: false,
      tooltip: {
        displayColors: false,
        callbacks: {
          title: toolTipTitle,
          label: toolTipLabel,
        },
      },
    },
  },
});
