from app.config.model_config import *
import replicate


def imagine(text):
    client = replicate.Client(api_token=REPLICATE_API_TOKEN)
    model = client.models.get(MODEL_PROTO)
    version = model.versions.get(MODEL_PROTO_VERSION)

    output = version.predict(prompt=text,
                             width=WIDTH,
                             height=HEIGHT,
                             prompt_strength=PROMPT_STRENGTH,
                             num_inference_steps=NUM_INFERENCE_STEPS
                             )[0]

    return output
