// ===== CONFIGURACIÓN GLOBAL =====
const API_BASE_URL = '/api/v1';
const TOKEN_KEY = 'auth_token';
const REFRESH_TOKEN_KEY = 'refresh_token';
const USER_KEY = 'user_data';

// Estado global
let currentUser = null;
let sessionTimer = null;
let tokenExpiryTime = null;
let charts = {};

// ===== UTILIDADES =====
const $ = (selector) => document.querySelector(selector);
const $$ = (selector) => document.querySelectorAll(selector);

const formatCurrency = (amount, currency = 'USD') => {
    return new Intl.NumberFormat('es-MX', {
        style: 'currency',
        currency: currency
    }).format(amount);
};

const formatDate = (dateString) => {
    const date = new Date(dateString);
    return date.toLocaleDateString('es-MX', {
        year: 'numeric',
        month: 'short',
        day: 'numeric'
    });
};

const showNotification = (message, type = 'info') => {
    const notification = document.createElement('div');
    notification.className = \`notification notification-\${type}\`;
    notification.textContent = message;
    notification.style.cssText = \`
        position: fixed;
        top: 80px;
        right: 20px;
        padding: 1rem 1.5rem;
        background: \${type === 'success' ? 'var(--success)' : type === 'error' ? 'var(--danger)' : 'var(--info)'};
        color: white;
        border-radius: 0.5rem;
        box-shadow: var(--shadow-md);
        z-index: 9999;
        animation: slideIn 0.3s ease;
    \`;
    document.body.appendChild(notification);
    setTimeout(() => {
        notification.style.animation = 'slideOut 0.3s ease';
        setTimeout(() => notification.remove(), 300);
    }, 3000);
};

// ===== API CLIENT =====
const api = {
    async request(endpoint, options = {}) {
        const token = localStorage.getItem(TOKEN_KEY);
        const headers = {
            'Content-Type': 'application/json',
            ...options.headers
        };
        if (token && !options.skipAuth) {
            headers['Authorization'] = `Bearer ${token}`;
        }
        const response = await fetch(`${API_BASE_URL}${endpoint}`, { ...options, headers });
        if (response.status === 401 && !options.skipAuth) {
            await this.refreshToken();
            return this.request(endpoint, options);
        }
        return await response.json();
    },
    async refreshToken() {
        const refreshToken = localStorage.getItem(REFRESH_TOKEN_KEY);
        if (!refreshToken) return false;
        try {
            const data = await this.request('/auth/refresh', {
                method: 'POST',
                skipAuth: true,
                body: JSON.stringify({ refreshToken })
            });
            localStorage.setItem(TOKEN_KEY, data.accessToken);
            updateTokenExpiry();
            return true;
        } catch (error) {
            return false;
        }
    },
    login: (data) => api.request('/auth/login', { method: 'POST', skipAuth: true, body: JSON.stringify(data) }),
    register: (data) => api.request('/auth/register', { method: 'POST', skipAuth: true, body: JSON.stringify(data) }),
    getProfile: () => api.request('/perfil'),
    updateProfile: (data) => api.request('/perfil', { method: 'PUT', body: JSON.stringify(data) }),
    getCategories: () => api.request('/categorias'),
    createCategory: (data) => api.request('/categorias', { method: 'POST', body: JSON.stringify(data) }),
    getTransactions: () => api.request('/transacciones'),
    createTransaction: (data) => api.request('/transacciones', { method: 'POST', body: JSON.stringify(data) }),
    deleteTransaction: (id) => api.request(`/transacciones/${id}`, { method: 'DELETE' }),
    getStatistics: (year, month) => api.request(`/reportes/estadisticas?year=${year}&month=${month}`),
    getUsers: () => api.request('/admin/usuarios'),
    approveUser: (id) => api.request(`/admin/usuarios/${id}/aprobar`, { method: 'PATCH' }),
    activateUser: (id) => api.request(`/admin/usuarios/${id}/activar`, { method: 'PATCH' }),
    deactivateUser: (id) => api.request(`/admin/usuarios/${id}/desactivar`, { method: 'PATCH' }),
    deleteUser: (id) => api.request(`/admin/usuarios/${id}`, { method: 'DELETE' })
};

// ===== AUTENTICACIÓN =====
function checkAuth() {
    const token = localStorage.getItem(TOKEN_KEY);
    const userData = localStorage.getItem(USER_KEY);
    if (token && userData) {
        currentUser = JSON.parse(userData);
        showDashboard();
    } else {
        showLogin();
    }
}

function updateTokenExpiry() {
    tokenExpiryTime = Date.now() + (15 * 60 * 1000);
    startSessionTimer();
}

function startSessionTimer() {
    if (sessionTimer) clearInterval(sessionTimer);
    sessionTimer = setInterval(() => {
        const remaining = tokenExpiryTime - Date.now();
        const minutes = Math.floor(remaining / 60000);
        const seconds = Math.floor((remaining % 60000) / 1000);
        if (remaining <= 0) {
            clearInterval(sessionTimer);
            logout();
            return;
        }
        const timerDisplay = $('#timerDisplay');
        if (timerDisplay) {
            timerDisplay.textContent = `${minutes}:${seconds.toString().padStart(2, '0')}`;
        }
        if (remaining <= 60000 && remaining > 59000) {
            showRefreshModal();
        }
    }, 1000);
}

function showRefreshModal() {
    const modal = $('#refreshModal');
    modal.classList.add('show');
    let countdown = 60;
    const countdownEl = $('#countdown');
    const countdownInterval = setInterval(() => {
        countdown--;
        countdownEl.textContent = countdown;
        if (countdown <= 0) {
            clearInterval(countdownInterval);
            logout();
        }
    }, 1000);
    $('#refreshBtn').onclick = async () => {
        clearInterval(countdownInterval);
        modal.classList.remove('show');
        await api.refreshToken();
        showNotification('Sesión renovada', 'success');
    };
    $('#logoutBtn').onclick = () => {
        clearInterval(countdownInterval);
        modal.classList.remove('show');
        logout();
    };
}

function showLogin() {
    $('#loginScreen').classList.add('active');
    $('#dashboardScreen').classList.remove('active');
}

function showDashboard() {
    $('#loginScreen').classList.remove('active');
    $('#dashboardScreen').classList.add('active');
    if (currentUser && currentUser.rol === 'admin') {
        document.body.classList.add('is-admin');
    }
    loadDashboardData();
}

function logout() {
    localStorage.clear();
    currentUser = null;
    if (sessionTimer) clearInterval(sessionTimer);
    showLogin();
}

// ===== EVENTOS DE LOGIN =====
function initLoginEvents() {
    $$('.login-tabs .tab-btn').forEach(btn => {
        btn.addEventListener('click', () => {
            $$('.login-tabs .tab-btn').forEach(b => b.classList.remove('active'));
            $$('.auth-form').forEach(f => f.classList.remove('active'));
            btn.classList.add('active');
            $(`#${btn.dataset.tab}Form`).classList.add('active');
        });
    });
    $$('.toggle-password').forEach(btn => {
        btn.addEventListener('click', () => {
            const input = btn.parentElement.querySelector('input');
            const icon = btn.querySelector('i');
            if (input.type === 'password') {
                input.type = 'text';
                icon.classList.replace('fa-eye', 'fa-eye-slash');
            } else {
                input.type = 'password';
                icon.classList.replace('fa-eye-slash', 'fa-eye');
            }
        });
    });
    $('#loginForm').addEventListener('submit', async (e) => {
        e.preventDefault();
        try {
            const data = await api.login({
                email: $('#loginEmail').value,
                password: $('#loginPassword').value
            });
            localStorage.setItem(TOKEN_KEY, data.accessToken);
            localStorage.setItem(REFRESH_TOKEN_KEY, data.refreshToken);
            currentUser = data.usuario;
            localStorage.setItem(USER_KEY, JSON.stringify(data.usuario));
            updateTokenExpiry();
            showDashboard();
            showNotification('Bienvenido!', 'success');
        } catch (error) {
            showNotification('Error al iniciar sesión', 'error');
        }
    });
    $('#registerForm').addEventListener('submit', async (e) => {
        e.preventDefault();
        const password = $('#registerPassword').value;
        const passwordConfirm = $('#registerPasswordConfirm').value;
        if (password !== passwordConfirm) {
            showNotification('Las contraseñas no coinciden', 'error');
            return;
        }
        try {
            await api.register({
                nombre: $('#registerNombre').value,
                email: $('#registerEmail').value,
                password: password
            });
            showNotification('Registro exitoso. Espera aprobación del admin.', 'success');
        } catch (error) {
            showNotification('Error al registrarse', 'error');
        }
    });
    $('#googleLoginBtn').addEventListener('click', () => {
        window.location.href = `${API_BASE_URL}/auth/google`;
    });
    $('#logoutLink').addEventListener('click', (e) => {
        e.preventDefault();
        logout();
    });
}

// ===== DASHBOARD =====
async function loadDashboardData() {
    try {
        const now = new Date();
        const [transactions, categories, stats] = await Promise.all([
            api.getTransactions(),
            api.getCategories(),
            api.getStatistics(now.getFullYear(), now.getMonth() + 1)
        ]);
        $('#userName').textContent = currentUser?.nombre || 'Usuario';
        $('#ingresosTotal').textContent = formatCurrency(stats.totalIngresos || 0);
        $('#egresosTotal').textContent = formatCurrency(stats.totalEgresos || 0);
        $('#balanceTotal').textContent = formatCurrency((stats.totalIngresos || 0) - (stats.totalEgresos || 0));
        updateCharts(stats);
        displayRecentTransactions(transactions.slice(0, 10));
    } catch (error) {
        console.error('Error:', error);
    }
}

function updateCharts(stats) {
    const ctx1 = $('#ingresosEgresosChart');
    if (ctx1) {
        if (charts.ingresosEgresos) charts.ingresosEgresos.destroy();
        charts.ingresosEgresos = new Chart(ctx1, {
            type: 'bar',
            data: {
                labels: ['Ingresos', 'Egresos'],
                datasets: [{
                    label: 'Monto',
                    data: [stats.totalIngresos || 0, stats.totalEgresos || 0],
                    backgroundColor: ['#a8e6cf', '#ffaaa5']
                }]
            },
            options: { responsive: true, plugins: { legend: { display: false } } }
        });
    }
}

function displayRecentTransactions(transactions) {
    const container = $('#recentTransactionsList');
    if (!container) return;
    container.innerHTML = transactions.map(t => `
        <div class="transaction-item">
            <div class="transaction-info">
                <div class="transaction-icon ${t.tipo}">
                    <i class="fas fa-arrow-${t.tipo === 'ingreso' ? 'up' : 'down'}"></i>
                </div>
                <div class="transaction-details">
                    <h4>${t.descripcion || 'Sin descripción'}</h4>
                    <p>${formatDate(t.fecha)}</p>
                </div>
            </div>
            <div class="transaction-amount ${t.tipo}">
                ${t.tipo === 'ingreso' ? '+' : '-'}${formatCurrency(t.monto, t.moneda)}
            </div>
        </div>
    `).join('');
}

// ===== NAVEGACIÓN =====
function initNavigation() {
    $$('.nav-item').forEach(item => {
        item.addEventListener('click', (e) => {
            e.preventDefault();
            $$('.nav-item').forEach(i => i.classList.remove('active'));
            item.classList.add('active');
            $$('.view').forEach(v => v.classList.remove('active'));
            $(`#${item.dataset.view}View`).classList.add('active');
            if (item.dataset.view === 'dashboard') loadDashboardData();
        });
    });
    $('#menuToggle')?.addEventListener('click', () => {
        $('#sidebar').classList.toggle('open');
    });
}

// ===== INICIALIZACIÓN =====
document.addEventListener('DOMContentLoaded', () => {
    const urlParams = new URLSearchParams(window.location.search);
    const token = urlParams.get('token');
    const refreshToken = urlParams.get('refreshToken');
    if (token && refreshToken) {
        localStorage.setItem(TOKEN_KEY, token);
        localStorage.setItem(REFRESH_TOKEN_KEY, refreshToken);
        updateTokenExpiry();
        window.history.replaceState({}, document.title, window.location.pathname);
    }
    initLoginEvents();
    initNavigation();
    checkAuth();
});
