import sys
import numpy as np
import sklearn.metrics as sm
import matplotlib.pyplot as plt
import pickle as pickle
from sklearn import linear_model


filename = sys.argv[1]
X = []
Y = []
with open(filename, 'r') as f:
    for line in f.readlines():
        xt, yt = [float(i) for i in line.split(',')]
        X.append(xt)
        Y.append(yt)

# train/test split
num_training = int(0.8 * len(X))
num_test = len(X) - num_training

# training data
X_train = np.array(X[:num_training]).reshape((num_training, 1))
Y_train = np.array(Y[:num_training])

# test data
X_test = np.array(X[num_training:]).reshape((num_test, 1))
Y_test = np.array(Y[num_training:])

# create linear regression object
linear_regressor = linear_model.LinearRegression()

# train model using the datasets
linear_regressor.fit(X_train, Y_train)

# predict the trained data output
y_train_pred = linear_regressor.predict(X_train)
plt.figure()
plt.scatter(X_train, Y_train, color='green')
plt.plot(X_train, y_train_pred, color='black', linewidth=4)
plt.title('Training Data')
plt.show()

Y_test_pred = linear_regressor.predict(X_test)
plt.figure()
plt.scatter(X_test, Y_test, color='green')
plt.plot(X_test, Y_test_pred, color='black', linewidth=4)
plt.title('Test Data')
plt.xticks(())
plt.yticks(())
plt.show()


# Measure performance
print('Mean absolute error =', round(sm.mean_absolute_error(Y_test, Y_test_pred), 2))
print('Mean squared error =', round(sm.mean_squared_error(Y_test, Y_test_pred), 2))
print('Median absolute error =', round(sm.median_absolute_error(Y_test, Y_test_pred), 2))
print('Explain variance score =', round(sm.explained_variance_score(Y_test, Y_test_pred), 2))
print('R2 score =', round(sm.r2_score(Y_test, Y_test_pred), 2))

# model persistance
output_model_file = '3_model_linear_regr.pkl'

with open(output_model_file, 'wb') as f:
    pickle.dump(linear_regressor, f)

with open(output_model_file, 'rb') as f:
    model_linregr = pickle.load(f)

y_test_pred_new = model_linregr.predict(X_test)
print('\nNew meand absolute error = ', round(sm.mean_absolute_error(Y_test, y_test_pred_new), 2))
