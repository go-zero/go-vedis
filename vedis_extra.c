#include "vedis_extra.h"

void vedis_error_message(vedis *store, const char **message)
{
    vedis_config(store, VEDIS_CONFIG_ERR_LOG, message, 0);
}
