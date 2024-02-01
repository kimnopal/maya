import * as jose from "jose";

export const getAuth = async (): Promise<object | null> => {
  let token;
  if (typeof window === "undefined") {
    const test = await import("next/headers");
    token = test.cookies().get("token")?.value;
  } else {
    const jsCookies = await require("js-cookie");
    token = jsCookies.get("token");
  }

  if (!token) {
    return null;
  }

  try {
    const decoded = await jose.jwtVerify(token, Buffer.from("rahasia", "utf8"));
    return decoded.payload;
  } catch (error) {
    return null;
  }
};
