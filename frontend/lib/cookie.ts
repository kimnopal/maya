import {cookies} from "next/headers"

export const setCookie = (data: string) => {
    cookies().set("auth", data, {
        maxAge: 60 * 60 * 48,
      });
}