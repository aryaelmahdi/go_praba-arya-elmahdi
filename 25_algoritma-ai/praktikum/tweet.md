# pengelompokkan tweet

- Pertama, pra-pemrosesan data dilakukan dengan memcah tweet menjadi kata-kata kemudian menghilangkan karakter khusus dan mengubah ke huruf kecil. 

- Kedua, ekstraksi fitur dari setiap tweet sebagai vektor fitur menggunakan metode seperti Word Embeddings. 
Langkah terakhir adalah memilih dan melatih model, dapat berupa model klasifikasi seperti model berbasis Transformer seperti BERT atau Bidirectional Encoder Representations from Transformers. model ini sangat baik karena dapat memahami hubungan antar kata-kata dan konteks dalam teks. 

- Setelah melatih model, klasifikasikan tweet baru berdasarkan sentimen dengan model tersebut. Hasil klasifikasi dapat digunakan untuk menganalisis mengelompokkan tweet berdasarkan sentimen.