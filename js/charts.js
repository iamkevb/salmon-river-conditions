function createTemperatureChart(id, labels, data) {
  const chartData = {
    labels: labels,
    datasets: [
      {
        backgroundColor: "rgba(54, 162, 235, 0.4)",
        borderColor: "rgba(54, 162, 235, 1)",
        hoverBackgroundColor: "rgba(54, 162, 235, 0.9)",
        borderWidth: 2,
        borderSkipped: false,
        data: data,
      },
    ],
  };

  const toolTipTitle = (toolTipItems) => {
    return toolTipItems[0].label;
  };

  const toolTipLabel = (toolTipItem) => {
    return [`High: ${toolTipItem.raw[1]}`, `Low: ${toolTipItem.raw[0]}`];
  };

  const config = {
    type: "bar",
    data: chartData,
    options: {
      maintainAspectRatio: false,
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
  }

  const ctx = document.getElementById(id);
  new Chart(ctx, config);
}

function createPrecipitationChart(id, labels, data) {
    console.log("PRECIP: ", id, data)
  const chartData = {
    labels: labels,
    datasets: [
      {
        label: "Rain",
        backgroundColor: "rgba(54, 162, 235, 0.4)",
        borderColor: "rgba(54, 162, 235, 1)",
        hoverBackgroundColor: "rgba(54, 162, 235, 0.9)",
        borderWidth: 2,
        data: data.rain,
      },
      {
        label: "Snow",
        backgroundColor: "rgba(255, 99, 132, 0.4)",
        borderColor: "rgba(255, 99, 132, 1)",
        hoverBackgroundColor: "rgba(255, 99, 132, 0.9)",
        borderWidth: 2,
        data: data.snow,
      },
    ],
  };

  const precipationToolTipTitle = (toolTipItems) => {
    return toolTipItems[0].label;
  };

  const precipationToolTipLabel = (toolTipItem) => {
    pType = toolTipItem.datasetIndex === 0 ? "Rain" : "Snow";
    pUnit = toolTipItem.datasetIndex === 0 ? "mm" : "cm";
    return `${pType}: ${toolTipItem.formattedValue} ${pUnit}`;
  };

  const config = {
    type: "bar",
    data: chartData,
    options: {
      maintainAspectRatio: false,
      plugins: {
        legend: false,
        tooltip: {
          displayColors: false,
          callbacks: {
            title: precipationToolTipTitle,
            label: precipationToolTipLabel,
          },
        },
      },
    },
  };

  const precipationCtx = document.getElementById(id);
  new Chart(precipationCtx, config);
}
