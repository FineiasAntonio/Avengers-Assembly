export function mostarNotificacao(message, type = 'success', duration = 3000) {
  const toast = document.createElement('div');
  toast.className = `notificacao ${type}`;
  toast.textContent = message;
  
  document.body.appendChild(toast);
  
  setTimeout(() => toast.classList.add('show'), 100);
  
  setTimeout(() => {
    toast.classList.remove('show');
    setTimeout(() => document.body.removeChild(toast), 300);
  }, duration);
}
