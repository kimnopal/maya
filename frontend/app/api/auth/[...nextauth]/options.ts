import CredentialsProvider from "next-auth/providers/credentials";
import bcrypt from "bcrypt";

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
        const user = {
          id: "1",
          name: "test",
          email: "testing@example.com",
          password:
            "$2a$12$duBa4bHZBQlU1dcZNEUVH.zLUr6HHsLRs3TgdJ/78CJg0heCgqVqm",
          majors: {
            name: "Teknik Elektro",
            faculty: "Teknik",
          },
        };

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
    signIn: "/profile/signin",
    signOut: "/profile/signout",
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
