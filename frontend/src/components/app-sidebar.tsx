import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarHeader,
} from "@/components/ui/sidebar";

/*
interface Artist {
  id: string;
  image: string;
  name: string;
}

interface Playlist {
  id: string;
  image: string;
  name: string;
}

interface Album {
  id: string;
  image: string;
  name: string;
  release_date: string;
}

const artists: Artist[] = [];
const playlists: Playlist[] = [];
const albums: Album[] = [];
*/

export function AppSidebar() {
  return (
    <Sidebar>
      <SidebarHeader />
      <SidebarContent>
        <SidebarGroup />
        <SidebarGroup />
      </SidebarContent>
      <SidebarFooter />
    </Sidebar>
  );
}
