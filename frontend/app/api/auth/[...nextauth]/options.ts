import CredentialsProvider from "next-auth/providers/credentials";
import bcrypt from "bcrypt";
import { findUser } from "@/app/action/action";
import { userType } from "@/app/types/type";

export const options = {
  providers: [
    CredentialsProvider({
      name: "Credentials",
      credentials: {
        email: {
          label: "email:",
          type: "text",
          placeholder: "your-email",
        },
        password: {
          label: "password:",
          type: "password",
          placeholder: "your-password",
        },
      },
      async authorize(credentials: any, request) {
        // Add logic here to look up the user from the credentials supplied
        const user: any = findUser({
          email: credentials.email,
        });

        if (user) {
          const match = await bcrypt.compare(
            credentials.password,
            user.password
          );

          if (match) {
            return user;
          }

          return null;
        } else {
          // If you return null then an error will be displayed advising the user to check their details.

          return null;

          // You can also Reject this callback with an Error thus the user will be sent to the error page with the error message as a query parameter
        }
      },
    }),
  ],
  pages: {
    signIn: "/signin",
    signOut: "/signout",
  },
  session: {
    jwt: true,
    maxAge: 60,
  },
  callbacks: {
    async jwt({ token, user }: { token: any; user: any }) {
      if (user) {
        token.majors = user.majors;
      }

      return token;
    },
    async session({ session, token }: { session: any; token: any }) {
      if (session?.user) {
        session.user.majors = token.majors;
      }

      return session;
    },
  },
};
