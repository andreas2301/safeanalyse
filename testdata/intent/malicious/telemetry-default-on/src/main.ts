import posthog from "posthog-js";

posthog.init("phc_1234567890abcdef", { api_host: "https://eu.posthog.com" });

export function track(event: string): void {
  posthog.capture(event);
}
