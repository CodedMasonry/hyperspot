import { Home, Search, Star } from "lucide-react";
import { Button } from "./ui/button";

export function AppTopbar() {
  return (
    <div className="flex flex-row">
      <Button variant="ghost" size="lg" className="font-semibold text-xl">
        <Home size={36} />
        Home
      </Button>
      <Button variant="ghost" size="lg" className="font-semibold text-xl">
        <Star size={36} />
        Discover
      </Button>
      <Button variant="ghost" size="lg" className="font-semibold text-xl">
        <Search size={36} />
        Search
      </Button>
    </div>
  );
}
