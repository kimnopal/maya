import React from "react";

const AuthTemplate = ({ children }: { children: React.ReactNode }) => {
  return <main className="flex flex-col gap-4 px-8 py-10">{children}</main>;
};

export default AuthTemplate;
