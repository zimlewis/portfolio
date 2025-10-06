document.addEventListener('alpine:init', () => {
    Alpine.data('navbar', () => ({
        open: false,
        toggleShow() {
            this.open = !this.open
        },
    }));


    Alpine.store('theme', {
        t: "light",
        set(theme) {
            this.t = theme
        },
        get() {
            return this.t
        }
    });
})
