import {
  type RouteConfig,
  route,
  index
} from "@react-router/dev/routes";

export default [
  index("App.tsx"),
  route("motoqueiro", "pages/Motoqueiro.tsx"),
  // * matches all URLs, the ? makes it optional so it will match / as well
  route("*?", "catchall.tsx"),
] satisfies RouteConfig;
