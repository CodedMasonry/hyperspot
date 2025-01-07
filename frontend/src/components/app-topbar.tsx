import { Home } from "lucide-react";
import { Button } from "./ui/button";

export function AppTopbar() {
  return (
    <div className="flex flex-row">
      <Button variant="ghost" size="lg" className="font-semibold text-2xl">
        <Home />
        Home
      </Button>
      <Button variant="ghost" size="lg">
        Discover
      </Button>
      <Button variant="ghost" size="lg">
        Search
      </Button>
    </div>
  );
}
