import type { AlpineComponent } from "alpinejs"

interface ThemeSelector {
    theme: string
    set: (theme: string) => void
    get: () => string
}


var selector : AlpineComponent<ThemeSelector> = {
    theme: "dark",
    set(t: string) {
        this.theme = t
    },
    get() {
        return this.theme
    }
}

export default selector;
