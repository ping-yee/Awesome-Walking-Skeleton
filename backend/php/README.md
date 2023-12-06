## Laravel + MySQL Walking Skeleton Solution

### Step to run

### 1. Install the dependency package.
```
composer install
```

### 2. Copy the env file.
```
cp .env.example .env
```

### 3. Generate the app key.
```
php artisan key:generate
```

### 4. Run the docker environment.
```
docker-compose up
```

### 5. Database Migration.
```
php artisan migrate
```

### 6. Seed your default data.
```
php artisan db:seed
```

### 7. Run the unit test.
```
./vendor/bin/phpunit
```