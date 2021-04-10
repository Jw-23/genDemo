import tensorflow_datasets as tfds

dataset, dataset_info = tfds.load('oxford_flowers102', data_dir='./src', split=',', with_info=True, as_supervised=True)
