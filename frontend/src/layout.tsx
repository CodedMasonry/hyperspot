import { SidebarProvider } from "@/components/ui/sidebar";
import { AppSidebar } from "@/components/app-sidebar";
import { AppTopbar } from "./components/app-topbar";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <SidebarProvider>
      <AppSidebar />
      <div className="mt-2 mx-2">
        <AppTopbar />
        <main className="m-2">{children}</main>
      </div>
    </SidebarProvider>
  );
}
