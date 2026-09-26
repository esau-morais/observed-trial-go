const status = document.getElementById('status');
const list = document.getElementById('deliveries');

async function loadRoute(route) {
  status.textContent = 'Loading…';
  const response = await fetch(`/api/deliveries?route=${route}`);
  const { deliveries } = await response.json();
  list.replaceChildren(
    ...deliveries.map((delivery) => {
      const item = document.createElement('li');
      item.textContent = `${delivery.id} · ${delivery.parcel} · ${delivery.status}`;
      return item;
    }),
  );
  status.textContent = `${deliveries.length} deliveries`;
}

for (const button of document.querySelectorAll('button[data-route]')) {
  button.addEventListener('click', () => loadRoute(button.dataset.route));
}
