import '../css/index.css';
import Alpine from 'alpinejs';
import selector from './alpines/theme';

window.Alpine = Alpine;

Alpine.store("theme", selector)

Alpine.start();
