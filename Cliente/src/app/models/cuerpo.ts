import {Tienda} from './tienda';

export interface Cuerpo {
  Datos: Indice[];
}

interface Indice {
  Indice: string;
  Departamentos: Departamento[];
}

interface Departamento {
  Nombre: string;
  Tiendas: Tienda[];
}
