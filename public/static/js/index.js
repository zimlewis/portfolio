document.addEventListener('alpine:init', () => {
    Alpine.data('navbar', () => ({
        open: false,
        toggleShow() {
            this.open = !this.open
        },
    }))
})
