export interface Calendario {
  Pedidos: Pedido[];
}

interface Pedido {
  Fecha: string;
  Tienda: string;
  Departamento: string;
  Calificacion: number;
  Productos: Producto[];
}

interface Producto {
  Codigo: number;
}
