export interface InventarioEntrada {
  Inventarios: Inventario[];
}

interface Inventario {
  Tienda: string;
  Departamento: string;
  Calificacion: number;
  Productos: Producto[];
}

interface Producto {
  Nombre: string;
  Codigo: number;
  Descripcion: string;
  Precio: number;
  Cantidad: number;
  Imagen: string;
}
