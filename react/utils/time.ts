export function unixToDate(timestamp: string) {
  return new Date(parseInt(timestamp) * 1000);
}

export function dateToHowLong(date: Date) {
  const now = new Date();
  // console.log("date", date);
  const diff = now.getTime() - date.getTime();
  const years = Math.floor(diff / 1000 / 60 / 60 / 24 / 365);
  const months = Math.floor(diff / (1000 * 60 * 60 * 24 * 30));
  const days = Math.floor(diff / (1000 * 60 * 60 * 24));
  const hours = Math.floor(diff / (1000 * 60 * 60));
  const minutes = Math.floor(diff / (1000 * 60));
  const seconds = Math.floor(diff / 1000);

  if (years > 0) {
    return `${years} y`;
  }

  if (months > 0) {
    return `${months} mo`;
  }

  if (days > 0) {
    return `${days}d`;
  }
  if (hours > 0) {
    return `${hours}h`;
  }

  if (minutes > 0) {
    return `${minutes}m`;
  }

  if (seconds > 0) {
    return `${seconds}s`;
  }
}
