# Hint: Implementar parseBody

## Paso 1: Buscar Content-Length

```go
clStr, ok := headers["Content-Length"]
if !ok {
    return nil, nil  // No hay body
}
```

## Paso 2: Convertir a número

```go
cl, err := strconv.Atoi(clStr)
if err != nil || cl <= 0 {
    return nil, nil
}
```

## Paso 3: Leer exactamente N bytes

```go
body := make([]byte, cl)
_, err = io.ReadFull(reader, body)
if err != nil {
    return nil, fmt.Errorf("error leyendo body: %w", err)
}
return body, nil
```

`io.ReadFull` garantiza que leemos exactamente `cl` bytes. Si el cliente envía menos, retorna error.
