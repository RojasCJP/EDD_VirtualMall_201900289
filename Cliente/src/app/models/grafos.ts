export interface Grafos {
  Nodos: Nodo[];
  PosicionInicialRobot: string;
  Entrega: string;
}

interface Enlace {
  Nombre: string;
  Distancia: number;
}

interface Nodo {
  Nombre: string;
  Enlaces: Enlace[];
}
