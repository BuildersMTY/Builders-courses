# Hint: Implementar parseRequestLine

## Paso 1: Leer la primera línea

```go
reqLine, err := reader.ReadString('\n')
if err != nil {
    return "", "", "", fmt.Errorf("error leyendo request line: %w", err)
}
```

`ReadString('\n')` lee hasta encontrar un salto de línea, incluyendo el `\n`.

## Paso 2: Limpiar y dividir

```go
reqLine = strings.TrimRight(reqLine, "\r\n")
parts := strings.Split(reqLine, " ")
if len(parts) != 3 {
    return "", "", "", fmt.Errorf("formato inválido: %q", reqLine)
}
```

## Paso 3: Retornar las partes

```go
return parts[0], parts[1], parts[2], nil
```

`parts[0]` es el método (GET, POST, etc.), `parts[1]` es el path, `parts[2]` es la versión HTTP.
