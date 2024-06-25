export const checkNull = (x: string | null | undefined): string => {
  if (x === null || x === "" || x === undefined) {
    return "";
  }
  return x;
};
