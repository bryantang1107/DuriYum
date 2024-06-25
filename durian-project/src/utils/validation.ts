const validationFunc = (
  funcName: string | undefined | null,
): ((x: string) => boolean) | undefined => {
  if (funcName === "password") {
    return function (password: string): boolean {
      //do validation
      return password.length >= 8;
    };
  }
  if (funcName === "email") {
    return function (email: string): boolean {
      //do validation
      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      return emailRegex.test(email);
    };
  }

  return undefined;
};

export default validationFunc;
