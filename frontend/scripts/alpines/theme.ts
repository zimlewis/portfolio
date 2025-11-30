import type { AlpineComponent } from "alpinejs"

type ThemePalette = "blue" | "green" | "orange" | "purple" | "red" | "rose" | "default" | "yellow"
type ThemeMode = "light" | "dark"

interface ThemeSelector {
    palette: ThemePalette
    mode: ThemeMode 
    setPalette: (palette: ThemePalette) => void
    setMode: (mode: ThemeMode) => void
    getPalette: () => ThemePalette
    getMode: () => ThemeMode 
}


var selector : AlpineComponent<ThemeSelector> = {
    palette: "default",
    mode: "dark",
    setPalette(palette: ThemePalette) {
        this.palette = palette;
    },
    setMode(mode: ThemeMode) {
        this.mode = mode;
    },
    getPalette() {
        return this.palette;
    },
    getMode() {
        return this.mode;
    }
}

export default selector;
