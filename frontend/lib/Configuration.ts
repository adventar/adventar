export function isDev() {
  return process.env.NODE_ENV === "development";
}

export function getToday(): Date {
  if (isDev() && process.env.CURRENT_DATE) {
    return new Date(process.env.CURRENT_DATE);
  } else {
    return new Date();
  }
}

export function getCurrentYear(): number {
  const today = getToday();
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
  const today = getToday();
  const month = today.getMonth() + 1;
  const day = today.getDate();
  return month === 11 || (month === 12 && day <= 25);
}
