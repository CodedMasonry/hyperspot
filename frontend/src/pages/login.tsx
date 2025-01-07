import { Button } from "@/components/ui/button";
import { AuthenticateUser } from "../../wailsjs/go/main/App";

export default function Login({
  setLoggedIn,
}: {
  setLoggedIn: React.Dispatch<React.SetStateAction<boolean>>;
}) {
  async function login() {
    // Tries to authenticate
    setLoggedIn(await AuthenticateUser());
  }

  return (
    <div className="flex flex-col items-center align-middle mt-36">
      <h1 className="font-bold text-3xl">Please Login Using Spotify.</h1>
      <Button onClick={login}>Login To Spotify</Button>
    </div>
  );
}
