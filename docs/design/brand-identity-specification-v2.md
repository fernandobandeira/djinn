# Djinn Brand Identity Specification v2.1 - Updated

## Brand Overview
Djinn is positioned as a premium AI assistant that grants wishes through intelligent automation. The brand balances approachability with sophistication, making AI feel magical yet trustworthy.

## Visual Identity

### Core Elements
**Primary Icon**: Magic lamp (teapot-style with spout), emerald green with blue/purple liquid inside, glass/ceramic quality
**Mascot**: The lamp itself (no separate character) - animates with liquid movement and magical effects
**Design Philosophy**: Clean elegance with strategic magic - pristine interfaces with purposeful liquid glass accents

### Color Palette
```
Primary Colors:
- Emerald Green: #10B981 (primary brand color)
- Light Emerald: #6EE7B7 (gradient complement)
- Magic Purple: #8B5CF6 (AI/magic elements)

Secondary Colors:
- Pure White: #FFFFFF (primary backgrounds - 70% of UI)
- Light Gray: #F8FAFC (subtle backgrounds)
- Text Gray: #374151 (primary text)

Accent Colors:
- Success Green: #10B981 (confirmations, success states)
- Premium Gold: #FCD34D (premium features, highlights)
- Error Red: #EF4444 (errors, warnings)

Gradients:
- Primary: Emerald Green (#10B981) → Light Emerald (#6EE7B7)
- Magic: Emerald Green → Magic Purple (#8B5CF6)
- Premium: Light Emerald → Premium Gold
```

### Typography
**Primary Font**: Inter (clean, modern sans-serif)
**Secondary Font**: Playfair Display (elegant serif for premium elements)
**Monospace**: JetBrains Mono (code and technical content)

### Logo Specifications
**Full Logo**: Magic lamp icon + "Djinn" wordmark
**Icon Only**: Standalone magic lamp symbol
**Wordmark Only**: Text-only version

**Lamp Icon Details**:
- Teapot-style silhouette with elegant spout
- Emerald green ceramic/glass material
- Blue/purple liquid visible inside
- Subtle glow effects for digital applications

**Minimum Sizes**:
- Full Logo: 120px width
- Icon Only: 24px
- Clear Space: 2x icon height on all sides

## Mascot Design

### Magic Lamp Mascot
**Personality**: The lamp itself embodies wisdom, helpfulness, and magical capability
**Visual Style**: Elegant ceramic/glass teapot form with animated liquid interior
**Key Features**:
- Emerald green ceramic/glass material
- Blue/purple magical liquid inside
- Spout releases gentle magical mist/sparkles
- Subtle glow effects around the base
- Liquid moves and swirls with user interactions

**Animation States**:
- Idle: Gentle liquid swirling, soft pulsing glow
- Thinking: Faster liquid movement, sparkles from spout
- Working: Dynamic liquid churning, brighter glow
- Success: Liquid settles into satisfied swirl, golden sparkles
- Error: Liquid slows, brief red tint, gentle shake
- Loading: Rhythmic liquid cycling, progressive glow

## Design System

### Component Hierarchy
1. **Primary**: Main actions, CTAs (emerald green)
2. **Secondary**: Supporting actions (light gray with emerald accents)
3. **Tertiary**: Subtle interactions (minimal styling)

### Visual Style Guidelines
**Balance Principle**: 70% solid UI, 20% gradients, 10% liquid glass accents

**Primary Style - Clean Solids**:
- Pure white backgrounds
- Solid color buttons and components
- Clear typography hierarchy
- Minimal shadows and borders

**Strategic Gradients** (20% of UI):
- Primary buttons and key CTAs
- Navigation highlights
- Success/completion states
- Premium feature indicators

**Liquid Glass Accents** (10% of UI):
- Lamp mascot and magical elements
- Loading states and transitions
- Hover effects on interactive elements
- Premium upgrade prompts

### Interaction Patterns
**Hover States**: Subtle emerald glow or gradient shift
**Click States**: Brief scale animation with color feedback
**Loading States**: Lamp liquid animation or progress gradients
**Success States**: Green confirmation with optional sparkle
**Error States**: Gentle shake with red accent, clear messaging

### Magical Element Usage
**Appropriate for**:
- Lamp mascot animations
- AI processing indicators
- Completion celebrations
- Premium feature highlights
- Onboarding moments

**Avoid for**:
- Basic navigation
- Form inputs
- Data tables
- Settings pages
- Error states (except gentle effects)

### Accessibility
- High contrast ratios (4.5:1 minimum)
- Clear focus indicators with emerald outlines
- Reduced motion options disable liquid animations
- Screen reader friendly with proper ARIA labels
- Keyboard navigation with visible focus states

## Brand Applications

### Digital Platforms
**Website**: Clean white layouts with emerald gradient CTAs and lamp mascot
**Mobile App**: Solid components with strategic gradients, animated lamp feedback
**Desktop App**: Expanded gradient usage, enhanced lamp animations
**Browser Extension**: Minimal lamp icon, emerald accent colors

### Marketing Materials
**Social Media**: Animated lamp mascot with magical liquid effects
**Documentation**: Clean white layouts with emerald headings and accents
**Presentations**: Professional white backgrounds with gradient highlights
**Email**: Simple lamp icon, emerald links, minimal magical elements

### User Experience Contexts
**Professional Users** (Sarah/Marcus/Robert):
- Emphasize clean, solid interfaces
- Minimal magical effects
- Focus on emerald accents for trust
- Lamp appears in success states

**Casual Users** (Alex/Zoe):
- More gradient usage allowed
- Enhanced lamp animations
- Playful magical feedback
- Richer visual celebrations

## Component Specifications

### Buttons
```css
/* Primary Button */
.btn-primary {
  background: linear-gradient(135deg, #10B981, #6EE7B7);
  color: white;
  border: none;
  border-radius: 8px;
  padding: 12px 24px;
  font-weight: 600;
  transition: all 0.2s ease;
}

.btn-primary:hover {
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
}

/* Secondary Button */
.btn-secondary {
  background: #F8FAFC;
  color: #374151;
  border: 1px solid #E5E7EB;
  border-radius: 8px;
  padding: 12px 24px;
  font-weight: 500;
}

.btn-secondary:hover {
  border-color: #10B981;
  color: #10B981;
}
```

### Cards
```css
.card {
  background: #FFFFFF;
  border: 1px solid #E5E7EB;
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  transition: all 0.2s ease;
}

.card:hover {
  border-color: #10B981;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.1);
}

/* Premium Card */
.card-premium {
  background: linear-gradient(135deg, #6EE7B7, #FCD34D);
  color: white;
  border: none;
}
```

### Form Elements
```css
.form-input {
  background: #FFFFFF;
  border: 1px solid #D1D5DB;
  border-radius: 8px;
  padding: 12px 16px;
  font-size: 16px;
  transition: border-color 0.2s ease;
}

.form-input:focus {
  outline: none;
  border-color: #10B981;
  box-shadow: 0 0 0 3px rgba(16, 185, 129, 0.1);
}

.form-input.error {
  border-color: #EF4444;
}

.form-input.success {
  border-color: #10B981;
}
```

### Navigation
```css
.nav-primary {
  background: #FFFFFF;
  border-bottom: 1px solid #E5E7EB;
  padding: 16px 24px;
}

.nav-link {
  color: #6B7280;
  text-decoration: none;
  padding: 8px 16px;
  border-radius: 6px;
  transition: all 0.2s ease;
}

.nav-link:hover {
  color: #10B981;
  background: rgba(16, 185, 129, 0.1);
}

.nav-link.active {
  color: #10B981;
  font-weight: 600;
  background: linear-gradient(135deg, rgba(16, 185, 129, 0.1), rgba(110, 231, 183, 0.1));
}
```

## Implementation Guidelines

### Design Token Structure
```json
{
  "colors": {
    "primary": {
      "emerald": "#10B981",
      "emerald-light": "#6EE7B7",
      "purple-magic": "#8B5CF6"
    },
    "neutral": {
      "white": "#FFFFFF",
      "gray-50": "#F8FAFC",
      "gray-700": "#374151"
    },
    "accent": {
      "success": "#10B981",
      "premium": "#FCD34D",
      "error": "#EF4444"
    }
  },
  "spacing": {
    "unit": "8px",
    "scale": ["4px", "8px", "12px", "16px", "24px", "32px", "48px", "64px"]
  },
  "typography": {
    "font-family": {
      "primary": "Inter, system-ui, sans-serif",
      "secondary": "Playfair Display, serif",
      "monospace": "JetBrains Mono, monospace"
    },
    "font-size": {
      "xs": "12px",
      "sm": "14px",
      "base": "16px",
      "lg": "18px",
      "xl": "20px",
      "2xl": "24px",
      "3xl": "32px"
    }
  }
}
```

### Animation Guidelines
```css
/* Lamp Mascot Animations */
@keyframes lamp-idle {
  0%, 100% { transform: rotate(-1deg); }
  50% { transform: rotate(1deg); }
}

@keyframes liquid-swirl {
  0% { transform: rotate(0deg) scale(1); }
  50% { transform: rotate(180deg) scale(1.1); }
  100% { transform: rotate(360deg) scale(1); }
}

@keyframes magical-glow {
  0%, 100% { box-shadow: 0 0 20px rgba(16, 185, 129, 0.3); }
  50% { box-shadow: 0 0 30px rgba(16, 185, 129, 0.5); }
}

/* Reduced Motion */
@media (prefers-reduced-motion: reduce) {
  .lamp-mascot,
  .magical-element {
    animation: none;
  }
}
```

### Responsive Behavior
**Mobile-First Approach**:
- Simplify magical effects on smaller screens
- Maintain emerald green brand recognition
- Ensure touch targets meet 44px minimum
- Reduce gradient complexity for performance

**Tablet Adaptations**:
- Enhanced lamp animations
- More strategic gradient usage
- Improved hover states

**Desktop Experience**:
- Full magical effect suite
- Rich lamp mascot interactions
- Advanced hover and focus states
- Premium glass morphism accents

## Quality Assurance

### Brand Consistency Checklist
- [ ] Emerald green (#10B981) used as primary brand color
- [ ] Lamp mascot represents brand (no separate genie character)
- [ ] 70/20/10 balance maintained (solid/gradient/glass)
- [ ] White backgrounds dominate interface
- [ ] Magical elements used purposefully, not everywhere
- [ ] Accessibility standards met (WCAG AA)
- [ ] Reduced motion options implemented
- [ ] Professional vs casual user contexts considered

### Performance Considerations
- Optimize gradient rendering for mobile devices
- Use CSS transforms for lamp animations
- Implement lazy loading for complex magical effects
- Provide fallbacks for older browsers
- Monitor animation frame rates

## Version History
- v2.1: Updated brand identity with lamp mascot, emerald green palette, balanced visual approach
- v2.0: Previous iteration with genie character
- v1.0: Initial brand specification with full glass morphism