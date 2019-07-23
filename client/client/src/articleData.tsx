export interface Article {
  id: number;
  title: string;
  content: string;
  imageNames: ImageName[];
}

interface ImageName {
  name: string;
}
