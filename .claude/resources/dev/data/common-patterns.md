# Common Implementation Patterns

## Design Patterns

### Creational Patterns

#### Singleton Pattern
```typescript
class ConfigManager {
  private static instance: ConfigManager;
  private config: Map<string, any>;
  
  private constructor() {
    this.config = new Map();
  }
  
  static getInstance(): ConfigManager {
    if (!ConfigManager.instance) {
      ConfigManager.instance = new ConfigManager();
    }
    return ConfigManager.instance;
  }
}
```

#### Factory Pattern
```typescript
interface Product {
  operation(): string;
}

class ProductFactory {
  static create(type: string): Product {
    switch (type) {
      case 'A': return new ProductA();
      case 'B': return new ProductB();
      default: throw new Error('Unknown product type');
    }
  }
}
```

#### Builder Pattern
```typescript
class QueryBuilder {
  private query: string = '';
  
  select(fields: string[]): this {
    this.query += `SELECT ${fields.join(', ')} `;
    return this;
  }
  
  from(table: string): this {
    this.query += `FROM ${table} `;
    return this;
  }
  
  where(condition: string): this {
    this.query += `WHERE ${condition} `;
    return this;
  }
  
  build(): string {
    return this.query.trim();
  }
}
```

### Structural Patterns

#### Repository Pattern
```typescript
interface Repository<T> {
  findById(id: string): Promise<T | null>;
  findAll(): Promise<T[]>;
  create(item: T): Promise<T>;
  update(id: string, item: Partial<T>): Promise<T>;
  delete(id: string): Promise<void>;
}

class UserRepository implements Repository<User> {
  async findById(id: string): Promise<User | null> {
    // Implementation
  }
  // Other methods...
}
```

#### Adapter Pattern
```typescript
// Third-party payment interface
interface PaymentGateway {
  makePayment(amount: number): void;
}

// Our interface
interface PaymentProcessor {
  processPayment(cents: number): boolean;
}

// Adapter
class PaymentAdapter implements PaymentProcessor {
  constructor(private gateway: PaymentGateway) {}
  
  processPayment(cents: number): boolean {
    try {
      this.gateway.makePayment(cents / 100);
      return true;
    } catch {
      return false;
    }
  }
}
```

### Behavioral Patterns

#### Observer Pattern
```typescript
interface Observer {
  update(data: any): void;
}

class Subject {
  private observers: Observer[] = [];
  
  attach(observer: Observer): void {
    this.observers.push(observer);
  }
  
  detach(observer: Observer): void {
    const index = this.observers.indexOf(observer);
    if (index > -1) {
      this.observers.splice(index, 1);
    }
  }
  
  notify(data: any): void {
    this.observers.forEach(observer => observer.update(data));
  }
}
```

#### Strategy Pattern
```typescript
interface SortStrategy {
  sort(data: number[]): number[];
}

class QuickSort implements SortStrategy {
  sort(data: number[]): number[] {
    // Quick sort implementation
  }
}

class MergeSort implements SortStrategy {
  sort(data: number[]): number[] {
    // Merge sort implementation
  }
}

class Sorter {
  constructor(private strategy: SortStrategy) {}
  
  setStrategy(strategy: SortStrategy): void {
    this.strategy = strategy;
  }
  
  sort(data: number[]): number[] {
    return this.strategy.sort(data);
  }
}
```

## State Management Patterns

### Redux Pattern
```typescript
// Action
const INCREMENT = 'INCREMENT';

// Action Creator
const increment = (amount: number) => ({
  type: INCREMENT,
  payload: amount
});

// Reducer
const counterReducer = (state = 0, action: any) => {
  switch (action.type) {
    case INCREMENT:
      return state + action.payload;
    default:
      return state;
  }
};
```

### Context Pattern (React)
```typescript
const ThemeContext = React.createContext<Theme>('light');

const ThemeProvider: React.FC = ({ children }) => {
  const [theme, setTheme] = useState<Theme>('light');
  
  return (
    <ThemeContext.Provider value={{ theme, setTheme }}>
      {children}
    </ThemeContext.Provider>
  );
};

const useTheme = () => {
  const context = useContext(ThemeContext);
  if (!context) {
    throw new Error('useTheme must be used within ThemeProvider');
  }
  return context;
};
```

## Error Handling Patterns

### Try-Catch-Finally
```typescript
async function riskyOperation(): Promise<Result> {
  let resource;
  try {
    resource = await acquireResource();
    return await processResource(resource);
  } catch (error) {
    logger.error('Operation failed', error);
    throw new ApplicationError('Processing failed', 'PROCESS_ERROR', 500);
  } finally {
    if (resource) {
      await releaseResource(resource);
    }
  }
}
```

### Result Pattern
```typescript
type Result<T, E = Error> = 
  | { success: true; value: T }
  | { success: false; error: E };

function divide(a: number, b: number): Result<number> {
  if (b === 0) {
    return { success: false, error: new Error('Division by zero') };
  }
  return { success: true, value: a / b };
}
```

### Circuit Breaker
```typescript
class CircuitBreaker {
  private failures = 0;
  private lastFailureTime?: Date;
  private state: 'CLOSED' | 'OPEN' | 'HALF_OPEN' = 'CLOSED';
  
  constructor(
    private threshold: number,
    private timeout: number
  ) {}
  
  async call<T>(fn: () => Promise<T>): Promise<T> {
    if (this.state === 'OPEN') {
      if (Date.now() - this.lastFailureTime!.getTime() > this.timeout) {
        this.state = 'HALF_OPEN';
      } else {
        throw new Error('Circuit breaker is OPEN');
      }
    }
    
    try {
      const result = await fn();
      this.onSuccess();
      return result;
    } catch (error) {
      this.onFailure();
      throw error;
    }
  }
  
  private onSuccess(): void {
    this.failures = 0;
    this.state = 'CLOSED';
  }
  
  private onFailure(): void {
    this.failures++;
    this.lastFailureTime = new Date();
    if (this.failures >= this.threshold) {
      this.state = 'OPEN';
    }
  }
}
```

## Async Patterns

### Promise Chain
```typescript
fetchUser(id)
  .then(user => fetchPosts(user.id))
  .then(posts => processPosts(posts))
  .then(result => console.log(result))
  .catch(error => console.error(error));
```

### Async/Await
```typescript
async function loadUserData(id: string): Promise<UserData> {
  try {
    const user = await fetchUser(id);
    const [posts, comments] = await Promise.all([
      fetchPosts(user.id),
      fetchComments(user.id)
    ]);
    return { user, posts, comments };
  } catch (error) {
    logger.error('Failed to load user data', error);
    throw error;
  }
}
```

### Debounce Pattern
```typescript
function debounce<T extends (...args: any[]) => any>(
  func: T,
  delay: number
): (...args: Parameters<T>) => void {
  let timeoutId: NodeJS.Timeout;
  
  return (...args: Parameters<T>) => {
    clearTimeout(timeoutId);
    timeoutId = setTimeout(() => func(...args), delay);
  };
}
```

## Caching Patterns

### Memoization
```typescript
function memoize<T extends (...args: any[]) => any>(fn: T): T {
  const cache = new Map();
  
  return ((...args: Parameters<T>) => {
    const key = JSON.stringify(args);
    if (cache.has(key)) {
      return cache.get(key);
    }
    const result = fn(...args);
    cache.set(key, result);
    return result;
  }) as T;
}
```

### LRU Cache
```typescript
class LRUCache<K, V> {
  private cache = new Map<K, V>();
  
  constructor(private maxSize: number) {}
  
  get(key: K): V | undefined {
    const value = this.cache.get(key);
    if (value !== undefined) {
      // Move to end (most recently used)
      this.cache.delete(key);
      this.cache.set(key, value);
    }
    return value;
  }
  
  set(key: K, value: V): void {
    if (this.cache.has(key)) {
      this.cache.delete(key);
    } else if (this.cache.size >= this.maxSize) {
      // Remove least recently used (first item)
      const firstKey = this.cache.keys().next().value;
      this.cache.delete(firstKey);
    }
    this.cache.set(key, value);
  }
}
```

## Testing Patterns

### Arrange-Act-Assert
```typescript
describe('UserService', () => {
  it('should create a new user', async () => {
    // Arrange
    const userData = { name: 'John', email: 'john@example.com' };
    const mockRepo = { save: jest.fn().mockResolvedValue({ id: 1, ...userData }) };
    const service = new UserService(mockRepo);
    
    // Act
    const user = await service.createUser(userData);
    
    // Assert
    expect(user).toHaveProperty('id');
    expect(user.name).toBe('John');
    expect(mockRepo.save).toHaveBeenCalledWith(userData);
  });
});
```

### Test Data Builder
```typescript
class UserBuilder {
  private user: Partial<User> = {};
  
  withName(name: string): this {
    this.user.name = name;
    return this;
  }
  
  withEmail(email: string): this {
    this.user.email = email;
    return this;
  }
  
  withAge(age: number): this {
    this.user.age = age;
    return this;
  }
  
  build(): User {
    return {
      id: '1',
      name: this.user.name || 'Default Name',
      email: this.user.email || 'default@example.com',
      age: this.user.age || 25
    };
  }
}
```