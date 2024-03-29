import SigninForm from "@/components/form/SigninForm";
import AuthTemplate from "@/components/template/AuthTemplate";
import Link from "next/link";

const LoginPage = () => {
  return (
    <AuthTemplate>
      <div className="flex justify-between items-end pb-6">
        <div className="flex flex-col gap-1">
          <h1 className="font-bold text-xl">Welcome to Maya</h1>
          <p className="text-xs">Sign in to get started!</p>
        </div>
        <Link
          href="/help"
          className="flex justify-center items-center gap-1 rounded-full shadow-md px-4 py-1"
        >
          <svg
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <g>
              <path
                d="M5.14282 12.0001V7.49157C5.15621 6.60533 5.34434 5.73047 5.69643 4.91705C6.04851 4.10366 6.55766 3.36773 7.19465 2.75143C7.83165 2.13514 8.58399 1.65059 9.40858 1.32556C10.2332 1.00053 11.1138 0.841396 12 0.857284C12.8861 0.841396 13.7668 1.00053 14.5913 1.32556C15.4159 1.65059 16.1683 2.13514 16.8053 2.75143C17.4423 3.36773 17.9515 4.10366 18.3036 4.91705C18.6555 5.73047 18.8437 6.60533 18.8571 7.49157V12.0001"
                stroke="#636F73"
              />
              <path
                d="M2.5714 9.42871H4.28569C4.51302 9.42871 4.73104 9.51902 4.89177 9.67977C5.05252 9.8405 5.14283 10.0585 5.14283 10.2859V15.4287C5.14283 15.656 5.05252 15.8741 4.89177 16.0348C4.73104 16.1955 4.51302 16.2859 4.28569 16.2859H2.5714C2.11674 16.2859 1.68071 16.1052 1.35922 15.7838C1.03773 15.4623 0.857117 15.0262 0.857117 14.5716V11.143C0.857117 10.6883 1.03773 10.2523 1.35922 9.93081C1.68071 9.60933 2.11674 9.42871 2.5714 9.42871Z"
                stroke="#636F73"
              />
              <path
                d="M21.4285 16.2859H19.7143C19.4869 16.2859 19.2689 16.1955 19.1081 16.0348C18.9475 15.8741 18.8571 15.656 18.8571 15.4287V10.2859C18.8571 10.0585 18.9475 9.8405 19.1081 9.67977C19.2689 9.51902 19.4869 9.42871 19.7143 9.42871H21.4285C21.8832 9.42871 22.3193 9.60933 22.6407 9.93081C22.9621 10.2523 23.1428 10.6883 23.1428 11.143V14.5716C23.1428 15.0262 22.9621 15.4623 22.6407 15.7838C22.3193 16.1052 21.8832 16.2859 21.4285 16.2859Z"
                stroke="#636F73"
              />
              <path
                d="M15.4285 21.0001C16.3378 21.0001 17.2098 20.6389 17.8529 19.9958C18.4959 19.3528 18.8571 18.4808 18.8571 17.5715V13.7144"
                stroke="#636F73"
              />
              <path
                d="M15.4286 21C15.4286 21.5683 15.2028 22.1135 14.8009 22.5153C14.3991 22.9171 13.854 23.1429 13.2857 23.1429H10.7143C10.1459 23.1429 9.60091 22.9171 9.19905 22.5153C8.79718 22.1135 8.57141 21.5683 8.57141 21C8.57141 20.4317 8.79718 19.8866 9.19905 19.4848C9.60091 19.0829 10.1459 18.8572 10.7143 18.8572H13.2857C13.854 18.8572 14.3991 19.0829 14.8009 19.4848C15.2028 19.8866 15.4286 20.4317 15.4286 21Z"
                stroke="#636F73"
              />
            </g>
          </svg>
          Help
        </Link>
      </div>
      <SigninForm />
      <p className="">
        Don't have an account?{" "}
        <Link
          href={{ pathname: "/profile/register" }}
          className="font-bold text-primary"
        >
          Register
        </Link>
      </p>
    </AuthTemplate>
  );
};

export default LoginPage;
