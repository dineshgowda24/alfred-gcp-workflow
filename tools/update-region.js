// to get the new list of regions and zones visit 
// https://cloud.google.com/compute/docs/regions-zones
// and run this script in the browser console
const rows = [...document.querySelectorAll("tbody.list tr")];
const regionMap = {};

rows.forEach(row => {
  const cells = row.querySelectorAll("td");
  const zone = cells[0]?.innerText.trim(); // e.g. "us-central1-a"
  const location = cells[1]?.innerText.trim();

  if (!zone || !location) return;

  const region = zone.split("-").slice(0, 2).join("-"); // e.g. "us-central1"

  if (!regionMap[region]) {
    regionMap[region] = {
      region,
      location,
      zones: []
    };
  }

  regionMap[region].zones.push(zone);
});

const result = Object.values(regionMap);
console.log(result);