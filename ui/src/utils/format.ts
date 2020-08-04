import { format } from "date-fns";

export const dateFormat = (date: string): string => {
  return format(new Date(date), "yyyy-MM-dd");
};

export const timeFormat = (time: string): string => {
  return format(new Date(time), "HH:mm");
};
