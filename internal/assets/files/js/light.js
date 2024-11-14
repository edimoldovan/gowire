class LightFramework {
  constructor(routes) {
    this.routes = routes;
    this.setupEventListeners();
  }

  setupEventListeners() {
    document.addEventListener("click", async (event) => {
      const target = event.target.closest("[data-url]");
      if (target) {
        event.preventDefault();
        const actionURL = target.getAttribute("data-url");
        const route = this.routes.find(r => r.path === actionURL);
        if (route) {
          await this.handleAction(route);
        }
      }
    });

    window.addEventListener("popstate", async (event) => {
      const route = this.routes.find(r => r.path === window.location.pathname);
      if (route) {
        await this.handleAction(route, false);
      }
    });
  }

  async handleAction(route, pushState = true) {
    try {
      const response = await fetch(route.path, {
        method: "GET",
        headers: {
          "X-Requested-With": "LightFramework"
        }
      });
      const text = await response.text();
      const parsed = new DOMParser().parseFromString(text, "text/html");
      document.querySelector("#content").replaceChildren(...parsed.querySelector("#content").children);

      if (pushState) {
        history.pushState(null, "", route.path);
      }

      document.title = `${route.path.charAt(1).toUpperCase() + route.path.slice(2)} - gowire`;
    } catch (error) {
      console.error("Error:", error);
    }
  }
}

// Initialize the framework with the routes
(async () => {
  try {
    const response = await fetch("/routes");
    const routes = await response.json();
    const light = new LightFramework(routes);
  } catch (error) {
    console.error("Error initializing LightFramework:", error);
  }
})();