export function isDev() {
  return process.env.NODE_ENV === "development";
}

export function getCurrentYear(): number {
  const today = new Date();
  const year = today.getFullYear();
  const month = today.getMonth() + 1;
  if (month < 11) {
    return year - 1;
  } else {
    return year;
  }
}

// 11/1 ~ 12/25
export function getCalendarCreatable(): boolean {
  if (isDev()) return true;
  const today = new Date();
  const month = today.getMonth() + 1;
  const day = today.getDate();
  return month === 11 || (month === 12 && day <= 25);
}
