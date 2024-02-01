import * as z from "zod";

export const LoginSchema = z.object({
  username: z.string(),
  password: z.string(),
});

export const RegisterSchema = z.object({
  name: z.string().min(1, {
    message: "Invalid name",
  }),
  username: z.string().min(1, {
    message: "Invalid username",
  }),
  email: z.string().email({
    message: "Invalid email",
  }),
  password: z.string().min(8, {
    message: "Invalid password",
  }),
});
