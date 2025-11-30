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

const getInitializePalette = (): ThemePalette => {
    if (localStorage.getItem('palette') as ThemePalette != null) {
        return localStorage.getItem('palette') as ThemePalette
    }
    return 'default'
}

const getInitializeMode = ():  ThemeMode => {
    if (localStorage.getItem('mode') as ThemeMode != null) {
        return localStorage.getItem('mode') as ThemeMode
    }

    if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
        return 'dark'
    }

    return 'light'
}


var selector : AlpineComponent<ThemeSelector> = {
    palette: getInitializePalette(),
    mode: getInitializeMode(),
    setPalette(palette: ThemePalette) {
        localStorage.setItem('palette', palette);
        this.palette = palette;
    },
    setMode(mode: ThemeMode) {
        localStorage.setItem('mode', mode);
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
