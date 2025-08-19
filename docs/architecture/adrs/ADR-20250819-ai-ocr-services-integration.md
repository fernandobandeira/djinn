# ADR-20250819: AI Assistant and OCR Services Integration Architecture

## Status
Accepted

## Context

Djinn's core value proposition relies on two critical AI-powered capabilities:
1. **OCR Receipt Itemization**: Converting receipt photos into structured data showing what users bought
2. **AI Financial Assistant**: Providing personalized financial advice based on spending patterns

### Business Requirements
- **Primary differentiator**: OCR receipt itemization with 90%+ accuracy
- **Privacy-first approach**: User data control and transparency
- **Freemium model**: Cost-effective processing for free tier users
- **Go backend + Flutter mobile**: Integration must align with existing architecture
- **Target accuracy**: 90%+ OCR accuracy for receipt text and line items

### Technical Constraints
- **Flutter integration**: Must work seamlessly within mobile app workflow
- **Offline capability**: Basic functionality when network unavailable
- **Cost optimization**: Minimize per-transaction costs for freemium model
- **Privacy compliance**: GDPR, CCPA compliance requirements
- **Performance**: Sub-3-second processing for good UX

### Current Challenges
- **OCR accuracy varies** greatly by receipt type (thermal paper, crumpled receipts, faded text)
- **Structured data extraction** beyond basic text recognition needed
- **Cost per receipt** ranges from $0.02-0.50 depending on solution
- **AI interaction costs** approximately $0.01 per conversation turn
- **Privacy vs accuracy trade-offs** between on-device and cloud processing

## Decision

### Cost-Optimized OCR Strategy for MVP

#### Primary OCR: Google ML Kit Text Recognition v2 (On-Device)
- **Reasoning**: Completely free with unlimited scans, no ongoing costs
- **Flutter integration**: Native `google_ml_kit` package support
- **Cost**: $0.00 per scan (on-device processing only)
- **Offline capability**: Full functionality without internet
- **Accuracy**: 90%+ for receipt text recognition
- **Privacy**: All processing happens on user's device

#### Future OCR: Tesseract (Alternative Free Option)
- **Purpose**: Backup for devices where ML Kit unavailable
- **Reasoning**: Completely free, open-source, on-device
- **Cost**: $0.00 per scan
- **Accuracy**: 85% (lower than ML Kit but still usable)
- **Integration**: Platform channels to native libraries

#### Post-MVP OCR: Cloud Services (When Scaling)
- **Options**: AWS Textract, Azure OCR, or Veryfi
- **Trigger**: Only when accuracy requirements exceed 95%
- **Cost**: $0.02-0.05 per scan (only for premium users)
- **Justification**: Premium feature for users needing highest accuracy

### Cost-Optimized AI Strategy for MVP

#### MVP Phase: DeepSeek R1 Free (OpenRouter)
- **Primary Model**: DeepSeek R1 (671B params, 37B active)
- **Capabilities**: 163K context, reasoning model, on par with GPT-4/Claude
- **Limits**: 50 req/day free (1000 req/day with $10 credits)
- **Cost**: $0.00 per token (just pay for higher rate limits)
- **Integration**: OpenRouter unified API
- **Quality**: State-of-the-art reasoning for financial advice

#### Production Phase: Premium Feature Strategy
- **Early Growth**: Continue DeepSeek R1 with purchased credits
- **Premium Tier**: AI advice as paid feature ($4.99/month)
- **Scale Options**: Upgrade to GPT-4o-mini or Claude Haiku if needed
- **Cost Structure**: ~$0.20-0.50/user for premium AI models
- **Margin**: 90-95% on premium subscriptions
- **Use case**: When Claude unavailable or specific features needed

### Architecture Integration

#### OCR Processing Pipeline
```
Receipt Image → Image Preprocessing → Google ML Kit
                                        ↓
                  Confidence < 85%? → AWS Textract
                                        ↓
                  Parse to structured data → Validate & Store
```

#### AI Assistant Flow
```
User Query + Receipt Data → Context Assembly → Claude API
                                                  ↓
                         Response Processing → Stream to Flutter UI
```

#### Flutter Integration Architecture
```dart
// OCR Service Abstraction
abstract class OCRService {
  Future<ReceiptData> processReceipt(File image);
}

// Multi-provider implementation
class HybridOCRService implements OCRService {
  final MLKitOCRProvider _primary;
  final TextractOCRProvider _fallback;
  
  @override
  Future<ReceiptData> processReceipt(File image) async {
    final result = await _primary.process(image);
    if (result.confidence < 0.85) {
      return await _fallback.process(image);
    }
    return result;
  }
}
```

### Data Privacy Implementation

#### On-Device Processing
- **Google ML Kit**: Processes images locally when possible
- **Receipt data**: Encrypted storage using Drift database encryption
- **AI context**: Limited to spending patterns, not raw receipt images

#### Cloud Processing Controls
- **AWS Textract**: Images deleted after processing
- **Claude API**: No conversation history retention beyond session
- **User consent**: Explicit opt-in for cloud processing features
- **Data deletion**: 30-day retention maximum

### Realistic Cost Progression Strategy

#### MVP Validation (10-50 users)
- **OCR**: Google ML Kit on-device ($0)
- **AI**: DeepSeek R1 free tier (50 req/day)
- **Total cost**: $0-10/month (just for rate limit upgrade)
- **Goal**: Validate AI adds value to users

#### Early Growth (100-500 users)
- **OCR**: Continue ML Kit on-device ($0)
- **AI**: DeepSeek R1 with credits (~20% of users active)
- **Revenue**: Premium tier at $4.99/month
- **Cost**: ~$20-50/month for API credits
- **Margin**: 90%+ on premium subscriptions

#### Production Scale (500+ users)
- **OCR**: ML Kit + cloud fallback for premium
- **AI**: Mix of DeepSeek R1 and premium models
- **Revenue**: $2,500+/month from premium users
- **Cost**: ~$100-200/month for APIs
- **Margin**: 92-96%

## Consequences

### Positive
- **High accuracy**: Multi-tier approach ensures 90%+ accuracy target
- **Cost effective**: Free tier coverage for majority of users
- **Privacy compliant**: On-device processing option available
- **Flutter native**: Seamless mobile integration
- **Scalable**: Architecture supports growth from freemium to enterprise
- **Service reliability**: Multiple providers prevent single point of failure
- **User choice**: Transparent privacy vs accuracy trade-offs

### Negative
- **Implementation complexity**: Multiple OCR providers increase codebase complexity
- **App size**: ML Kit and potential PaddleOCR increase binary size
- **Network dependency**: Cloud features require reliable internet
- **Vendor dependencies**: Reliance on Google, AWS, and Anthropic services
- **Cost monitoring**: Need sophisticated usage tracking for cost control

### Risks
- **API cost escalation**: Cloud providers may increase pricing
  - *Mitigation*: Multi-provider strategy allows switching
- **Privacy regulation changes**: GDPR/CCPA may restrict cloud processing
  - *Mitigation*: On-device PaddleOCR provides compliance path
- **Service availability**: Cloud APIs may have outages
  - *Mitigation*: Graceful degradation to offline-only features
- **Accuracy degradation**: Receipt formats may change over time
  - *Mitigation*: Regular accuracy monitoring and provider evaluation

## Alternatives Considered

### Option A: Single Cloud Provider (AWS Only)
- **Description**: Use AWS Textract for all OCR processing
- **Pros**: Consistent API, excellent structured data extraction
- **Cons**: Higher costs ($0.50/receipt), no offline capability, vendor lock-in
- **Reason for rejection**: Cost prohibitive for freemium model

### Option B: Pure On-Device Processing
- **Description**: Tesseract or PaddleOCR exclusively
- **Pros**: Zero API costs, complete privacy, offline capability
- **Cons**: Lower accuracy (~75-85%), larger app size, development complexity
- **Reason for rejection**: Cannot meet 90% accuracy requirement

### Option C: Receipt-Specific Services (Veryfi/Taggun)
- **Description**: Use specialized receipt processing APIs
- **Pros**: Highest accuracy (96-99%), built-in line-item extraction
- **Cons**: Very high cost ($0.20-0.50/receipt), vendor lock-in
- **Reason for rejection**: Unsustainable for freemium business model

### Option D: OpenAI GPT-4 Vision for OCR
- **Description**: Use GPT-4 Vision API for receipt processing
- **Pros**: Single API for OCR and AI, improving capabilities
- **Cons**: Higher latency, more expensive than dedicated OCR
- **Reason for rejection**: Cost and performance inferior to specialized OCR

## Implementation Notes

### Phase 1: Core Implementation (Weeks 1-4)
1. **Google ML Kit integration**: Flutter plugin setup for on-device OCR
2. **OpenRouter API integration**: GPT-OSS-20B free model setup
3. **Receipt data models**: Structured data representation and validation
4. **Basic UI**: Camera capture, processing indicators, results display

#### OpenRouter Configuration with DeepSeek R1
```dart
// OpenRouter client setup for DeepSeek R1
final openRouterClient = http.Client();
const openRouterUrl = 'https://openrouter.ai/api/v1/chat/completions';

Future<String> callAI(String prompt) async {
  final response = await openRouterClient.post(
    Uri.parse(openRouterUrl),
    headers: {
      'Authorization': 'Bearer \$OPENROUTER_API_KEY',
      'HTTP-Referer': 'https://djinn.app', // Required
      'X-Title': 'Djinn Finance App', // Optional
    },
    body: jsonEncode({
      'model': 'deepseek/deepseek-r1:free', // 671B reasoning model
      'messages': [{'role': 'user', 'content': prompt}],
      'temperature': 0.1, // Low for financial accuracy
      'max_tokens': 4096, // Reasonable limit for responses
    }),
  );
  return response.body;
}
```

### Phase 2: Enhanced Processing (Weeks 5-8)
1. **AWS Textract fallback**: Implement secondary OCR processing
2. **Image preprocessing**: Perspective correction, contrast enhancement
3. **Accuracy monitoring**: Confidence scoring and quality metrics
4. **Cost tracking**: Usage monitoring and billing integration

### Phase 3: Optimization (Weeks 9-12)
1. **PaddleOCR integration**: On-device processing option
2. **Caching layer**: Processed receipt storage and retrieval
3. **Offline queuing**: Process receipts when connectivity restored
4. **A/B testing**: Provider performance comparison

### Success Metrics
- **OCR accuracy**: >90% for receipt text recognition
- **Processing time**: <3 seconds end-to-end
- **Cost per receipt**: <$0.02 for premium users
- **User satisfaction**: >4.5 rating for OCR feature
- **Privacy compliance**: Zero data retention policy violations

### Monitoring Strategy
- **Accuracy tracking**: Regular evaluation against manual verification
- **Cost monitoring**: Daily usage and cost alerts
- **Performance metrics**: Processing time and failure rate tracking
- **User feedback**: In-app accuracy rating collection

### Rollback Plan
- **Provider switching**: Immediate fallback to alternative OCR providers
- **Feature degradation**: Disable cloud processing if privacy concerns arise
- **Manual processing**: Human verification workflow for critical receipts
- **Data recovery**: Export processed receipt data before provider changes

## References
- [Related ADR: Personal Finance Tech Stack](./ADR-20250812-personal-finance-tech-stack.md)
- [OCR Solutions Architecture Comparison](../OCR-Receipt-Solutions-Architecture-Comparison.md)
- [Google ML Kit Text Recognition Documentation](https://developers.google.com/ml-kit/vision/text-recognition/v2)
- [AWS Textract Expense Analysis](https://docs.aws.amazon.com/textract/latest/dg/analyzing-document-expense.html)
- [Anthropic Claude API Documentation](https://docs.anthropic.com/claude/reference)
- [Flutter ML Integration Guide](https://flutter.dev/docs/development/data-and-backend/ml)
- [GDPR AI Processing Guidelines](https://gdpr-info.eu/art-22-gdpr/)

## Decision Makers
- **Author**: System Architect
- **Reviewers**: Product Manager, Lead Developer, Privacy Officer
- **Approver**: Technical Director
- **Date**: 2025-08-19

## Revision History
- **2025-08-19**: Initial draft with multi-tier OCR and AI strategy