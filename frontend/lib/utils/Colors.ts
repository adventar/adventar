const colors = [
  "#FFA000",
  "#5942A4",
  "#00A98D",
  "#68B0C7",
  "#8DBF38",
  "#F70021",
  "#787875",
  "#317C8C",
  "#C94030",
  "#4A9769",
  "#CF9731",
  "#A3BE37",
  "#FA226C",
  "#004584",
  "#D8613F",
  "#8AC5D2",
  "#75A961"
];

export function calendarColor(calendarId: number): string {
  return colors[calendarId % colors.length];
}
